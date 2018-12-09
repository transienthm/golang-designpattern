package structure

import "fmt"

type Facade struct {
	M Music
	V Video
	C Count
}

func (this *Facade) GetRecommandVideos() error {
	this.V.GetVideos()
	return nil
}

type Video struct {
	vid int64
}

func (this *Video) GetVideos() error {
	fmt.Println("get videos")
	return nil
}

type Music struct {
}

func (this *Music) GetMusic() error {
	fmt.Println("get music material")
	return nil
}

type Count struct {
	PraiseCnt int64
	CommentCnt int64
	CollectCnt int64
}

func (this *Count) GetCountById(id int64) (*Count, error) {
	fmt.Println("get video counts")
	return this, nil
}

func main() {
	f := &Facade{}
	f.GetRecommandVideos()
}