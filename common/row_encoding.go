package common

import (
	"fmt"
)

// Values in rows are typically encoded in little-endian order
// Most CPU architectures are little-endian so those allows us to simply cast values in the cast of int types

func EncodeRow(row *Row, colTypes []ColumnType, buffer []byte) ([]byte, error) {
	for colIndex, colType := range colTypes {
		var err error
		buffer, err = encodeRowCol(row, colIndex, colType, buffer)
		if err != nil {
			return nil, err
		}
	}
	return buffer, nil
}

func encodeRowCol(row *Row, colIndex int, colType ColumnType, buffer []byte) ([]byte, error) {
	if row.IsNull(colIndex) {
		buffer = append(buffer, 0)
	} else {
		buffer = append(buffer, 1)
		switch colType.Type {
		case TypeTinyInt, TypeInt, TypeBigInt:
			// We store as unsigned so convert signed to unsigned
			valInt64 := row.GetInt64(colIndex)
			buffer = AppendUint64ToBufferLE(buffer, uint64(valInt64))
		case TypeDecimal:
			valDec := row.GetDecimal(colIndex)
			var err error
			buffer, err = valDec.Encode(buffer, colType.DecPrecision, colType.DecScale)
			if err != nil {
				return nil, err
			}
		case TypeDouble:
			valFloat64 := row.GetFloat64(colIndex)
			buffer = AppendFloat64ToBufferLE(buffer, valFloat64)
		case TypeVarchar:
			valString := row.GetString(colIndex)
			buffer = AppendStringToBufferLE(buffer, valString)
		default:
			return nil, fmt.Errorf("unexpected column type %d", colType)
		}
	}
	return buffer, nil
}

func DecodeRow(buffer []byte, colTypes []ColumnType, rows *Rows) error {
	offset := 0
	for colIndex, colType := range colTypes {
		if buffer[offset] == 0 {
			offset++
			rows.AppendNullToColumn(colIndex)
		} else {
			offset++
			switch colType.Type {
			case TypeTinyInt, TypeInt, TypeBigInt:
				var u uint64
				u, offset = ReadUint64FromBufferLE(buffer, offset)
				rows.AppendInt64ToColumn(colIndex, int64(u))
			case TypeDecimal:
				var val Decimal
				var err error
				val, offset, err = ReadDecimalFromBuffer(buffer, offset, colType.DecPrecision, colType.DecScale)
				if err != nil {
					return err
				}
				rows.AppendDecimalToColumn(colIndex, val)
			case TypeDouble:
				var val float64
				val, offset = ReadFloat64FromBufferLE(buffer, offset)
				rows.AppendFloat64ToColumn(colIndex, val)
			case TypeVarchar:
				var val string
				val, offset = ReadStringFromBuffer(buffer, offset)
				rows.AppendStringToColumn(colIndex, val)
			default:
				return fmt.Errorf("unexpected column type %d", colType)
			}
		}
	}
	return nil
}