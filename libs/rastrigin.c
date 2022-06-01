#include "rastrigin.h"
#include <math.h>

#ifndef M_PI
#define M_PI		3.14159265358979323846
#endif

double rastrigin(int n)
{
  const int A = 10;
  double res, x, inc;

  inc = 10.24 / n;
  x = -5.12;

  for (int i = 0; i < n; ++i)
  {
    res += (x * x) - A * cos(2.0 * M_PI * x);
    x += inc;
  }
  res = (A * n) + res;

  return res;
}
