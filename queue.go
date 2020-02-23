package main

//getSong
func (v *VoiceInstance) QueueGetSong() (song Song) {
	v.queueMutex.Lock()
	defer v.queueMutex.Unlock()
	if len(v.queue) != 0 {
		return v.queue[0]
	}
	return
}

//Add Song to queue
func (v *VoiceInstance) QueueAdd(song Song) {
	v.queueMutex.Lock()
	defer v.queueMutex.Unlock()
	v.queue = append(v.queue, song)
}

//Skip
func (v *VoiceInstance) QueueSkip() {
	v.queueMutex.Lock()
	defer v.queueMutex.Unlock()
	if len(v.queue) != 0 {
		v.queue = v.queue[1:]
	}
}

//Remove Song from index
func (v *VoiceInstance) QueueremoveIndex(index int) {
	v.queueMutex.Lock()
	defer v.queueMutex.Unlock()
	if len(v.queue) != 0 && index <= len(v.queue) {
		v.queue = append(v.queue[:index], v.queue[index+1:]...)
	}
}

//QueueRemoveUser
func (v *VoiceInstance) QueueRemoveUser(user string) {
	v.queueMutex.Lock()
	defer v.queueMutex.Unlock()
	queue := v.queue
	v.queue = []Song{}
	if len(v.queue) != 0 {
		for _, q := range queue {
			if q.User != user {
				v.queue = append(v.queue, q)
			}
		}
	}
}

//QueueRemoveLast
func (v *VoiceInstance) QueueRemoveLast() {
	v.queueMutex.Lock()
	defer v.queueMutex.Unlock()
	if len(v.queue) != 0 {
		v.queue = append(v.queue[:len(v.queue)-1], v.queue[len(v.queue)-1]...)
	}
}

//QueueClean
func (v *VoiceInstance) QueueClean() {
	v.queueMutex.Lock()
	defer v.queueMutex.Unlock()
	v.queue = v.queue[:1]
}

//QueueRemove
func (v *VoiceInstance) QueueRemove() {
	v.queueMutex.Lock()
	defer v.queueMutex.Unlock()
	v.queue = []Song{}
}
