
type Message struct {
	iStartTime int
	iEndTime   int
}

type AuthenticationManager struct {
	stMapData   map[string]*Message
	iTimeToLive int
}

func Constructor(timeToLive int) AuthenticationManager {
	obj := new(AuthenticationManager)
	obj.iTimeToLive = timeToLive
	obj.stMapData = make(map[string]*Message)
	return *obj
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	_, _ok := this.stMapData[tokenId]
	if _ok {
		return
	}

	pMessage := new(Message)
	pMessage.iStartTime = currentTime
	pMessage.iEndTime = currentTime + this.iTimeToLive
	this.stMapData[tokenId] = pMessage
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	pValue, _ok := this.stMapData[tokenId]
	if !_ok {
		return
	}

	if currentTime >= pValue.iEndTime{
		return
	}

	pValue.iEndTime = currentTime + this.iTimeToLive
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	iCnt := 0
	for _, v := range this.stMapData {
		if currentTime > v.iStartTime && currentTime < v.iEndTime {
			iCnt++
		}
	}
	return iCnt
}