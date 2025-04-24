package main
import "fmt"
func main(){
 var ano int
 fmt.Scan(&ano)
  for ano > 2015{
	ano = ano - 12
  }
  if ano == 2004 {
	fmt.Println("Macaco")
  } else if ano == 2005 {
	fmt.Println("Galo")
  } else if ano == 2006 {
	fmt.Println("Cao")
} else if ano == 2007 {
	fmt.Println("Porco")
  } else if ano == 2008 {
	fmt.Println("Rato")
  } else if ano == 2009 {
	fmt.Println("Boi")
  } else if ano == 2010 {
	fmt.Println("Tigre")
  } else if ano == 2011 {
	fmt.Println("Coelho")
  } else if ano == 2012 {
	fmt.Println("Dragao")
  } else if ano == 2013 {
	fmt.Println("Serpente")
  } else if ano == 2014 {
	fmt.Println("Cavalo")
  } else if ano == 2015 {
	fmt.Println("Cabra")
  }}