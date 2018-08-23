#include "MyDLL.h"  
  
extern "C" int CalcSum_(int a, int b, int c);

 int sj()
 {
	  int a = 17, b = 11, c = 14;
      int sum = CalcSum_(a, b, c);
	 return sum;
 }
  
