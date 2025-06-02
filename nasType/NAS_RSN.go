package nasType

// RSN 9.11.4.33
type RSN struct {
	Iei   uint8
	Len   uint8
	Octet uint8
}

func NewRSN(iei uint8) (rsn *RSN) {
	rsn = &RSN{}
	rsn.SetIei(iei)
	return rsn
}

func (a *RSN) GetIei() (iei uint8) {
	return a.Iei
}

func (a *RSN) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *RSN) GetLen() (len uint8) {
	return a.Len
}

func (a *RSN) SetLen(len uint8) {
	a.Len = len
}

func (a *RSN) GetRSN() uint8 {
	return a.Octet & GetBitMask(1, 0)
}

func (a *RSN) SetRSN(rsn uint8) {
	a.Octet = (a.Octet & 0xFE) + (rsn & 0x01)
}
