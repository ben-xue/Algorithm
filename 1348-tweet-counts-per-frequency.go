type SecondLinkNode struct {
	iTimeSecond int
	iCnt        int
	pNextNode   *SecondLinkNode
}

type MinuteLinkNode struct {
	iTimeMinute int
	iTweetCnt   int
	pNextNode   *SecondLinkNode
}

type UserTweetCounts struct{
	indexMap map[int] *MinuteLinkNode
}

type TweetCounts struct {
	nameMap map[string] *UserTweetCounts
}

type TimeSeqInSecond struct {
	iStart   int
	iEnd     int
	bFullSeq bool
}

func ConstructorUser() *UserTweetCounts {
	UserTweetCounts := *new(UserTweetCounts)
	UserTweetCounts.indexMap = make(map[int] *MinuteLinkNode)
	return &UserTweetCounts
}

func Constructor() TweetCounts {
	TweetCounts := *new(TweetCounts)
	TweetCounts.nameMap = make(map[string] *UserTweetCounts)
	return TweetCounts
}

var SecondPerMinute int = 60
var SecondPerHour int = 60 * 60
var SecondPerDay int = 60 * 60 * 24

func (this *TweetCounts) InsertSecondNode(pMinuteNode *MinuteLinkNode, iTime int) {
	pNewSecondNode := new(SecondLinkNode)
	pNewSecondNode.iCnt++
	pNewSecondNode.iTimeSecond = iTime
	pNewSecondNode.pNextNode = nil

	//find
	pTmpNode := pMinuteNode.pNextNode
	var pPrevNode *SecondLinkNode = nil
	for pTmpNode != nil && pTmpNode.iTimeSecond < pNewSecondNode.iTimeSecond {
		pPrevNode = pTmpNode
		pTmpNode = pTmpNode.pNextNode
	}

	//change list
	if pTmpNode != nil {
		pNewSecondNode.pNextNode = pTmpNode
		if pPrevNode != nil{
			pPrevNode.pNextNode = pNewSecondNode
		}
	} else {
		if pPrevNode == nil{
			pMinuteNode.pNextNode = pNewSecondNode
		}else{
			pPrevNode.pNextNode = pNewSecondNode
		}
	}

	pMinuteNode.iTweetCnt++
}

func (this *TweetCounts) RecordTweet(tweetName string, time int) {

	//find map by name
	pUser,_ok := this.nameMap[tweetName]
	if !_ok{
		pUser = ConstructorUser()
		this.nameMap[tweetName] = pUser
	}

	iMinute := time / SecondPerMinute
	pValue, _ok := pUser.indexMap[iMinute]
	if !_ok {
		//not found
		pNewMinteNode := new(MinuteLinkNode)
		pNewMinteNode.iTimeMinute = iMinute
		pNewMinteNode.pNextNode = nil

		this.InsertSecondNode(pNewMinteNode, time)

		pUser.indexMap[iMinute] = pNewMinteNode
	} else {
		this.InsertSecondNode(pValue, time)
	}
}

func (this *TweetCounts) GetTimeSeqInSecond(freq string, tweetName string, startTime int, endTime int) []TimeSeqInSecond {
	iInterval := 0
	switch freq {
	case "minute":
		{
			iInterval = SecondPerMinute
		}
	case "hour":
		{
			iInterval = SecondPerHour
		}
	case "day":
		{
			iInterval = SecondPerDay
		}
	}

	result := []TimeSeqInSecond{}

	for i := startTime; i <= endTime; i += iInterval {
		stTmpSeq := TimeSeqInSecond{0,0,false}
		if i + iInterval <= endTime{
			stTmpSeq = TimeSeqInSecond{i, i + iInterval, false}
		}else{
			stTmpSeq = TimeSeqInSecond{i, endTime + 1, false}
		}
		result = append(result, stTmpSeq)
	}
	return result
}

func (this *TweetCounts) GetTweetCntInTimeSeq(pUser *UserTweetCounts ,startTime int, endTime int) int {
	iStartMinute := startTime / SecondPerMinute
	iEndMinute := (endTime - 1)/ SecondPerMinute

	iCnt := 0
	for i := iStartMinute ; i <= iEndMinute ;i++ {
		_,_ok := pUser.indexMap[i]
		if !_ok{
			continue
		}
		iTmpCnt := 0
		for pTmp := pUser.indexMap[i].pNextNode; pTmp != nil; pTmp = pTmp.pNextNode {
			if pTmp.iTimeSecond >= startTime {
				iTmpCnt++
			}
		}
		iCnt += iTmpCnt
	}
	return iCnt
}

func (this *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
	result := []int{}

	pUser,_ok := this.nameMap[tweetName]
	if !_ok{
		fmt.Println("error")
		return nil
	}

	stTimeSeqList := this.GetTimeSeqInSecond(freq,tweetName,startTime,endTime)
	for i := range stTimeSeqList {
		iCnt := this.GetTweetCntInTimeSeq(pUser,stTimeSeqList[i].iStart,stTimeSeqList[i].iEnd)
		result = append(result,iCnt)
	}

	return result
}

