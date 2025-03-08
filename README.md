# Задача о бросании иголки
Имеется поле с параллельными прямыми, идущими через одно и тоже расстояние. Соседние параллельные прямые находятся на расстоянии a. Случайным образом в это поле бросается игла длины l.
Вопрос: какова вероятность, что игла пересечет хотя бы одну прямую?

![Рисунок иголки с обозначениями l, a, y, 𝛂](https://github.com/user-attachments/assets/a60beab7-258e-469c-a1ec-ea88826b79db)

Как работает алгоритм:
1. Сперва задаются длина между линиями _а_, длина иглы _l_, размер поля по _Х_ и _Y_, количество бросков.
2. Для каждого броска генерируется случайное значение координат середины иголки и угол под которым она падает _𝛂_.
3. Определяются координаты линий, находящихся в поле и заносятся в отдельный массив.
4. Определяется перпендикуляр от середины иглы _y_.
5. Находим, пересекает ли игла линию: проходим по каждому значению массива с координатами линий и смотрим, если координаты иглы по Y больше координат сравниваемой линии, то в случае, если перпендикуляр _y_ пересекает ближайшую линию слева или справа, засчитываем, что игла пересекла линию.

Пусть _x_ - расстояние от центра иглы до ближайшей прямой, тогда область всех возможных положений иглы можно записать следующийм образом:
  Ω = {0 ≤ 𝛂 ≤ π; 0 ≤ x ≤ a}
  
![График](https://github.com/user-attachments/assets/4fae04a4-c113-4134-9b30-64edbebb8d2a)

  |Ω| = π * a

Если p - относительная частота пересечения иглой линий, то:
  π = _l_/pa

Таким образом, используя последнюю формулу, мы можем определить, насколько точна полученная вероятность (чем ближе к числу π, тем точнее результат)
