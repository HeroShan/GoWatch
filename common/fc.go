package common

func InArray(need interface{}, needArr []interface{}) bool {
	for _,v := range needArr{
	   if need == v{
		   return true
	   }
   }
   return false
}