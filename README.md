
# МТРПЗ перша лабораторна

Цей проєкт призначений для розв’язання квадратних рівнянь виду:

a(x^2) + bx + c = 0

де (a!=0). Програма реалізована мовою Go та підтримує два режими запуску:  
1. Інтерактивний (без аргументів)  
2. Неінтерактивний (файловий) (з одним аргументом – шлях до файлу)  


---

## Короткий опис застосунку

- Зчитує коефіцієнти (a, b, c).
- Перевіряє, чи (a!=0). Якщо (a = 0), виводить помилку, оскільки це не квадратне рівняння.
- Обчислює дискримінант (d = b^2 - 4ac).
- Залежно від знака дискримінанта виводить 0, 1 або 2 дійсні корені.
- Виводить рівняння та знайдені корені (якщо вони існують).

## Інструкція: як зібрати та запустити проєкт

1. Клонувати репозиторій (якщо ще не клоновано):
   git clone https://github.com/MikhailoSafronov/software_first_lab.git
3. Перейти до папки проєкту
4. Зібрати (скомпілювати) проєкт
go build -o equation main.go
У результаті з’явиться виконуваний файл equation (на Windows – equation.exe).
5. Запустити програму
Інтерактивний режим (без аргументів):
./equation
Програма попросить у консолі ввести a,b,c.
Неінтерактивний (файловий) режим (з одним аргументом):
./equation test_valid.txt
Де test_valid.txt – файл, що містить принаймні три числа (коефіцієнти a,b,c).
## Формат файлу для неінтерактивного режиму
Файл для неінтерактивного режиму повинен містити три числа, що відповідають коефіцієнтам 
a,b,c. Вони можуть бути розділені пробілами або символами нового рядка.
У разі відсутності потрібних чисел або якщо вказані нечислові дані, програма виведе повідомлення про помилку і завершиться з ненульовим кодом.
## Вказівка на revert-коміт

В історії комітів цього репозиторію присутній коміт із назвою:Revert "Added information to Readme.md file"
Це демонструє вміння відкотити (revert) зміни, які були додані у попередньому коміті, без переписування історії. Ви можете переглянути цей коміт у списку (через git log або на GitHub) і побачити, як він анулює зміни з попереднього коміту.
