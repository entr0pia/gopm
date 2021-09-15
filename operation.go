package main

func (t *urlBase) search(more bool) {
	err := t.query()
	if more {
		for true {
			t.page += 1
			err = t.query()
			if err != nil {
				break
			}
		}
	}
}
