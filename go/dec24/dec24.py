from z3 import *

px = Int('px')
py = Int('py')
pz = Int('pz')
vx = Int('vx')
vy = Int('vy')
vz = Int('vz')
t0 = Int('t0')
t1 = Int('t1')
t2 = Int('t2')

s = Solver()
s.add(px + vx *t0 == 258040640626680 + 46 * t0)
s.add(py + vy *t0 == 312821058997749 -15*t0)
s.add(pz + vz *t0 == 253745803620360 + 29 * t0)
s.add(px + vx *t1 == 342350112377290 - 42 * t1)
s.add(py + vy *t1 == 337453812074005 + 57 * t1)
s.add(pz + vz *t1 == 216597774328568 - 128 * t1)
s.add(px + vx *t2 == 272589395905592 - 13 * t2)
s.add(py + vy *t2 == 362308351555047 - 116 * t2)
s.add(pz + vz *t2 == 281416220790402 + 15 * t2)
s.add(t0 > 0, t1 > 0, t2 > 0)
print(s.check())
print(s.model())

# Answer from Z3:
#
#[pz = 156592420220258,
# vx = -164,
# py = 368909610239045,
# vz = 223,
# px = 363206674204110,
# t1 = 170955424810,
# t0 = 500790636083,
# vy = -127,
# t2 = 600114425818]
