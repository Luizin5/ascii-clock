package clock

import (
   "github.com/mbndr/figlet4go"
   tm "github.com/inancgumus/screen"
   "time"
)
func Start(loop bool) {
   render := figlet4go.NewAsciiRender()
   t := time.Now()

   clock, _ := render.Render(t.Format("15:04"))

   if loop {
      for {
	 t = time.Now()
	 tm.Clear()
	 clock,_ = render.Render(t.Format("15:04"))
	 println(clock)
	 tm.MoveTopLeft()
	 time.Sleep(time.Second)
      }
   }else {
      println(clock)
   }
}
