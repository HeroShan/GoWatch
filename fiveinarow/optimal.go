package fiveinarow

import(
	"math"
)

func (all *Allinat)Center()  {
	var c Coordinat
	if len(all.Key) == 0{
		c.X = int(math.Ceil(Vertical/2))
		c.Y = int(math.Ceil(Level/2))
		all.AddCoordinat(c)
	}else{
		c.X = all.Key[0].X+1
		c.Y = all.Key[0].Y
		all.AddCoordinat(c)
	}
	
	return
}