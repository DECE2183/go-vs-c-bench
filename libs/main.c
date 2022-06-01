
#include <stdio.h>
#include <time.h>

#include "rastrigin.h"

int main(void)
{
  clock_t start_time, elapsed_time;

  int n = (1 << 24);

  start_time = clock();
  double res = rastrigin(n);
  elapsed_time = clock() - start_time;

  printf("C result: %f, time: %ld ms\n", res, elapsed_time);
  return 0;
}
