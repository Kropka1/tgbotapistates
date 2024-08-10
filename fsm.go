package tgbotapi

func (u User) SetState(stateDestiny string) {
	u.State = State{
		Destiny: stateDestiny,
	}
}

func (u User) GetState() {

}
