// Code generated by github.com/spacemeshos/go-scale/scalegen. DO NOT EDIT.

// nolint
package types

import (
	"github.com/spacemeshos/go-scale"
)

func (t *Transaction) EncodeScale(enc *scale.Encoder) (total int, err error) {
	{
		n, err := t.RawTx.EncodeScale(enc)
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeOption(enc, t.TxHeader)
		if err != nil {
			return total, err
		}
		total += n
	}
	return total, nil
}

func (t *Transaction) DecodeScale(dec *scale.Decoder) (total int, err error) {
	{
		n, err := t.RawTx.DecodeScale(dec)
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		field, n, err := scale.DecodeOption[TxHeader](dec)
		if err != nil {
			return total, err
		}
		total += n
		t.TxHeader = field
	}
	return total, nil
}

func (t *Reward) EncodeScale(enc *scale.Encoder) (total int, err error) {
	{
		n, err := t.Layer.EncodeScale(enc)
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeCompact64(enc, uint64(t.TotalReward))
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeCompact64(enc, uint64(t.LayerReward))
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeByteArray(enc, t.Coinbase[:])
		if err != nil {
			return total, err
		}
		total += n
	}
	return total, nil
}

func (t *Reward) DecodeScale(dec *scale.Decoder) (total int, err error) {
	{
		n, err := t.Layer.DecodeScale(dec)
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		field, n, err := scale.DecodeCompact64(dec)
		if err != nil {
			return total, err
		}
		total += n
		t.TotalReward = uint64(field)
	}
	{
		field, n, err := scale.DecodeCompact64(dec)
		if err != nil {
			return total, err
		}
		total += n
		t.LayerReward = uint64(field)
	}
	{
		n, err := scale.DecodeByteArray(dec, t.Coinbase[:])
		if err != nil {
			return total, err
		}
		total += n
	}
	return total, nil
}

func (t *RawTx) EncodeScale(enc *scale.Encoder) (total int, err error) {
	{
		n, err := scale.EncodeByteArray(enc, t.ID[:])
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeByteSlice(enc, t.Raw)
		if err != nil {
			return total, err
		}
		total += n
	}
	return total, nil
}

func (t *RawTx) DecodeScale(dec *scale.Decoder) (total int, err error) {
	{
		n, err := scale.DecodeByteArray(dec, t.ID[:])
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		field, n, err := scale.DecodeByteSlice(dec)
		if err != nil {
			return total, err
		}
		total += n
		t.Raw = field
	}
	return total, nil
}