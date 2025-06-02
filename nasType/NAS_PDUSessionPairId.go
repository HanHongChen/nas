package nasType

// PDUSessionPairId 9.11.4.32
type PDUSessionPairId struct {
	Iei   uint8
	Len   uint8
	Octet uint8
}

func NewPDUSessionPairId(iei uint8) (pduSessionPairId *PDUSessionPairId) {
	pduSessionPairId = &PDUSessionPairId{}
	pduSessionPairId.SetIei(iei)
	return pduSessionPairId
}

func (a *PDUSessionPairId) GetIei() (iei uint8) {
	return a.Iei
}

func (a *PDUSessionPairId) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *PDUSessionPairId) GetLen() (len uint8) {
	return a.Len
}

func (a *PDUSessionPairId) SetLen(len uint8) {
	a.Len = len
}

func (a *PDUSessionPairId) GetPduSessionId() uint8 {
	return a.Octet & GetBitMask(3, 0)
}

func (a *PDUSessionPairId) SetPduSessionId(id uint8) {
	a.Octet = (a.Octet & 0xF8) + (id & 0x07)
}
