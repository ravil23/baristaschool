package collection

import (
	"github.com/ravil23/baristaschool/telegrambot/entity"
)

var Quiz = entity.NewQuiz(map[entity.Question]entity.Answer{
    //  
    "Кофейное дерево опыляется": {
        CorrectOption: "Самоопылением",
        InvalidOptions: []string{"Насекомыми", "Ветром", "Водой"},
    },
    "Кофейное дерево — вечнозеленое растение семейства ...": {
        CorrectOption: "Мареновые",
        InvalidOptions: []string{"Яблоневые", "Коммелиновые", "Кутровые"},
    },
    "Кофейный пояс находится между тропиками": {
        CorrectOption: "Рака и Козерога",
        InvalidOptions: []string{"Водолея и Стрельца", "Рака и Водолея", "Стрельца и Козерога"},
    },
    "Период цветения кофейного дерева": {
        CorrectOption: "3 дня",
        InvalidOptions: []string{"День", "Неделя", "Месяц"},
    },
    "Реакция Майяра происходит между": {
        CorrectOption: "Сахарами и кислотами",
        InvalidOptions: []string{"Сахарами", "Кислотами", "Ни одно из перечисленных"},
    },
    "Слово 'робуста' переводится как ...": {
        CorrectOption: "Сильный",
        InvalidOptions: []string{"Горький", "Плохой", "Быстрый"},
    },

    //  Виды кофе
    "Арабика появилась в результате естественного скрещивания ... и ...": {
        CorrectOption: "Робуста и эужениодис",
        InvalidOptions: []string{"Робуста и либерика", "Эксцельса и либерика", "Робуста и эксцельса"},
    },
    "Коммерческую ценность представляют": {
        CorrectOption: "Арабика и робуста",
        InvalidOptions: []string{"Арабика и эксельса", "Эужениоидис и робуста", "Либерийка и экцельса"},
    },

    //  Диапазон цифр
    "...% кофеина в арабике": {
        CorrectOption: "0,8-1,5",
        InvalidOptions: []string{"0,5-1", "0,8-1", "1-1,5"},
    },
    "Арабика произрастает на высоте ...м": {
        CorrectOption: "8-2300",
        InvalidOptions: []string{"0-800", "0-1000", "800-1000"},
    },
    "В робусте содержится ...% кофеина": {
        CorrectOption: "1,6-3,5",
        InvalidOptions: []string{"0,8-1,5", "0-0,8", "1,9-2,3"},
    },
    "Дерево арабики вырастает в высоту ... м": {
        CorrectOption: "5-10",
        InvalidOptions: []string{"3-10", "0,5-15", "5-15"},
    },
    "Дерево арабики даёт урожай ... раза/год": {
        CorrectOption: "1-2",
        InvalidOptions: []string{"2-4", "1-3", "3-4"},
    },
    "Диапазон высоты произрастания кофейного дерева": {
        CorrectOption: "0-2300",
        InvalidOptions: []string{"0-800", "800-1600", "800-2300"},
    },
    "Период созревания ягоды робусты составляет ... месяцев": {
        CorrectOption: "9-12",
        InvalidOptions: []string{"3-6", "6-12", "3-9"},
    },
    "Робуста произрастает на высоте ... м": {
        CorrectOption: "0-800",
        InvalidOptions: []string{"800-2300", "800-1000", "0-1000"},
    },
    "Средняя температура произрастания кофе +...°-...°": {
        CorrectOption: "15-30",
        InvalidOptions: []string{"10-25", "10-30", "15-25"},
    },
    "Ферментация кофейной ягоды в мытом обработке длится ... ч": {
        CorrectOption: "12-72",
        InvalidOptions: []string{"6-12", "8-10", "80-100"},
    },

    //  Доля
    "Всего арабика составляет ... мирового производства": {
        CorrectOption: "2/3",
        InvalidOptions: []string{"1/3", "1/2", "1/4"},
    },

    //  Кофемолка для альтернативы
    "#1 Что из перечисленного является кофемолкой для альтернативы": {
        CorrectOption: "Ножевая",
        InvalidOptions: []string{"Кофемолка с контролем времени", "Кофемолка с ручным дозатором", "Кофемолка с контролем веса"},
    },
    "#2 Что из перечисленного является кофемолкой для альтернативы": {
        CorrectOption: "Ручная",
        InvalidOptions: []string{"Кофемолка с контролем времени", "Кофемолка с ручным дозатором", "Кофемолка с контролем веса"},
    },
    "#3 Что из перечисленного является кофемолкой для альтернативы": {
        CorrectOption: "Производственная",
        InvalidOptions: []string{"Кофемолка с контролем времени", "Кофемолка с ручным дозатором", "Кофемолка с контролем веса"},
    },
    "#4 Что из перечисленного является кофемолкой для альтернативы": {
        CorrectOption: "Лабораторная",
        InvalidOptions: []string{"Кофемолка с контролем времени", "Кофемолка с ручным дозатором", "Кофемолка с контролем веса"},
    },
    "#5 Что из перечисленного является кофемолкой для альтернативы": {
        CorrectOption: "Ни одно из перечисленного",
        InvalidOptions: []string{"Кофемолка с контролем времени", "Кофемолка с ручным дозатором", "Кофемолка с контролем веса"},
    },
    "#6 Что из перечисленного является кофемолкой для альтернативы": {
        CorrectOption: "Все перечисленные",
        InvalidOptions: []string{"Лабораторная", "Ручная", "Производственная"},
    },

    //  Кофемолка для эспрессо
    "#1 Что из перечисленного является кофемолкой для эспрессо": {
        CorrectOption: "Бункерная",
        InvalidOptions: []string{"Лабораторная", "Производственная", "Ножевая"},
    },
    "#2 Что из перечисленного является кофемолкой для эспрессо": {
        CorrectOption: "Кофемолка с контролем времени",
        InvalidOptions: []string{"Ручная", "Ножевая", "Производственная"},
    },
    "#3 Что из перечисленного является кофемолкой для эспрессо": {
        CorrectOption: "Ни одно из перечисленного",
        InvalidOptions: []string{"Лабораторная", "Ручная", "Производственная"},
    },
    "#4 Что из перечисленного является кофемолкой для эспрессо": {
        CorrectOption: "Кофемолка с контролем времени",
        InvalidOptions: []string{"Ножевая", "Лабораторная", "Ручная"},
    },
    "#5 Что из перечисленного является кофемолкой для эспрессо": {
        CorrectOption: "Все перечисленные",
        InvalidOptions: []string{"Кофемолка с контролем времени", "Кофемолка с ручным дозатором", "Кофемолка с контролем веса"},
    },

    //  Обозначения
    "Зерно средней твёрдости (700-900м)": {
        CorrectOption: "MHB",
        InvalidOptions: []string{"PB", "SHB", "HB"},
    },
    "Очень твёрдое зерно (>1500)": {
        CorrectOption: "SHB",
        InvalidOptions: []string{"GBN", "HB", "MHB"},
    },
    "Плотное зерно (0-700м)": {
        CorrectOption: "PB",
        InvalidOptions: []string{"SHB", "HB", "MHB"},
    },
    "Твёрдое зерно (900-1200м)": {
        CorrectOption: "HB",
        InvalidOptions: []string{"MNB", "PB", "SHB"},
    },
    "Хорошее твёрдое зерно (1200-1500м)": {
        CorrectOption: "GHB",
        InvalidOptions: []string{"HB", "MNB", "PB"},
    },

    //  Пена
    "Количество пены для американо": {
        CorrectOption: "0",
        InvalidOptions: []string{"1", "0,5", "1,5"},
    },
    "Количество пены капучино ... см": {
        CorrectOption: "1,5",
        InvalidOptions: []string{"1", "2-3", "0,5"},
    },
    "Количество пены латте": {
        CorrectOption: "1",
        InvalidOptions: []string{"0,5", "2-3", "1,5"},
    },
    "Количество пены латте-маккиато": {
        CorrectOption: "2-3",
        InvalidOptions: []string{"0,5", "1", "1,5"},
    },
    "Количество пены раф кофе": {
        CorrectOption: "1,5",
        InvalidOptions: []string{"1", "0,5", "2-3"},
    },
    "Количество пены флэт уайта": {
        CorrectOption: "0,5",
        InvalidOptions: []string{"2-3", "1,5", "1"},
    },

    //  Питчеры
    "Питчер объемом ... мл используется для взбивания молока для напитков 300-500 мл": {
        CorrectOption: "600",
        InvalidOptions: []string{"150", "350", "1000"},
    },
    "Питчер объемом ... мл используется для взбивания молока для напитков объемом 200-250 мл": {
        CorrectOption: "350",
        InvalidOptions: []string{"150", "600", "1000"},
    },
    "Питчер объемом ... со используется для приготовления импрессионист в э/м с низкими рабочими группами": {
        CorrectOption: "150",
        InvalidOptions: []string{"600", "350", "1000"},
    },

    //  Послевкусие
    "Послевкусие идеально экстрагированного кофе": {
        CorrectOption: "Длительное",
        InvalidOptions: []string{"Быстрое", "Пустое"},
    },
    "Послевкусие недоэкстрагированного кофе": {
        CorrectOption: "Быстрое",
        InvalidOptions: []string{"Длительное", "Пустое"},
    },
    "Послевкусие переэкстрагированного кофе": {
        CorrectOption: "Пустое",
        InvalidOptions: []string{"Быстрое", "Длительное"},
    },

    //  Разновидность
    "Первая естественная мутация арабики типики": {
        CorrectOption: "Бурбон",
        InvalidOptions: []string{"Гейша", "Иргалем", "Сеосиес"},
    },
    "Самая известная разновидность арабики": {
        CorrectOption: "Гейша",
        InvalidOptions: []string{"Бурбон", "Иргалем", "Сеосиес"},
    },

    //  Страна
    "Родина арабики": {
        CorrectOption: "Эфиопия",
        InvalidOptions: []string{"Индия", "Кения", "Вьетнам"},
    },

    //  Цифры
    "Всего видов Coffea": {
        CorrectOption: "120",
        InvalidOptions: []string{"67", "100", "134"},
    },
    "Дерево робусты вырастает в высоту ... м": {
        CorrectOption: "15",
        InvalidOptions: []string{"5", "10", "20"},
    },
    "Для продолжения рода робусты нужно ... дерева за счёт перекрестного опыления": {
        CorrectOption: "2",
        InvalidOptions: []string{"1", "3", "4"},
    },
    "Количество хромосом арабики": {
        CorrectOption: "44",
        InvalidOptions: []string{"12", "22", "48"},
    },
    "Количество хромосом робусты": {
        CorrectOption: "22",
        InvalidOptions: []string{"6", "12", "44"},
    },
    "Кофейное дерево даёт урожай ... лет": {
        CorrectOption: "25",
        InvalidOptions: []string{"10", "15", "20"},
    },
    "Кофейное дерево начинает плодоносить спустя ... года после посадки": {
        CorrectOption: "3",
        InvalidOptions: []string{"1", "2", "4"},
    },
    "Кофейный пояс расположен в ...° от экватора в двух направлениях": {
        CorrectOption: "25",
        InvalidOptions: []string{"10", "15", "20"},
    },
    "Продолжительность жизни кофейного дерева — ... лет": {
        CorrectOption: "50",
        InvalidOptions: []string{"15", "25", "35"},
    },
    "Робуста была признана видом кофе в ... веке": {
        CorrectOption: "19",
        InvalidOptions: []string{"9", "16", "13"},
    },
    "Сушка кофейной ягоды происходит до ...% влажности": {
        CorrectOption: "11",
        InvalidOptions: []string{"9", "20", "42"},
    },

    //  Экстракция
    "Идеально экстрагированный кофе — вкус ...": {
        CorrectOption: "Сладкий, спелый и понятный, богатая кислотность",
        InvalidOptions: []string{"Кислый, нехватка сладости, солёный", "Горький, сухой, вяжущий"},
    },
    "Недоэкстракт — вкус ...": {
        CorrectOption: "Кислый, нехватка сладости, солёный",
        InvalidOptions: []string{"Сладкий, спелый и понятный, богатая кислотность", "Горький, сухой, вяжущий"},
    },
    "Переэкстракт — вкус": {
        CorrectOption: "Горький, сухой, вяжущий",
        InvalidOptions: []string{"Кислый, нехватка сладости, солёный", "Сладкий, спелый и понятный, богатая кислотность"},
    },

    //  Этапы обжарки
    "Второй этап обжарки": {
        CorrectOption: "Высушивание",
        InvalidOptions: []string{"Первый крек", "Второй крек", "Карамелизация"},
    },
    "Данная реакция придаёт кофе коричневый цвет": {
        CorrectOption: "Карамелизация",
        InvalidOptions: []string{"Второй крек", "Высушивание", "Первый крек"},
    },
    "На этой стадии обжарки выделяется запах травянистого оттенка": {
        CorrectOption: "Высушивание",
        InvalidOptions: []string{"Гомогенизация", "Реакция Майяра", "Второй крек"},
    },
    "Первый этап обжарки кофе": {
        CorrectOption: "Гомогенизация",
        InvalidOptions: []string{"Высушивание", "Реакция Майяра", "Первый крек"},
    },
    "Приводит к появлению масел на поверхности зерна, кислотность отсутствует, много горечи.": {
        CorrectOption: "Второй крек",
        InvalidOptions: []string{"Первый крек", "Развитие", "Высушивание"},
    },
    "Пятый этап обжарки": {
        CorrectOption: "Первый крек",
        InvalidOptions: []string{"Развитие", "Второй крек", "Карамелизация"},
    },
    "Седьмой этап обжарки": {
        CorrectOption: "Второй крек",
        InvalidOptions: []string{"Развитие", "Первый крек", "Реакция Майяра"},
    },
    "Температура 150° характерна для этапа обжарки": {
        CorrectOption: "Реакция Майяра",
        InvalidOptions: []string{"Первый крек", "Высушивание", "Второй крек"},
    },
    "Температура 160-180° характерна для этого этапа обжарки": {
        CorrectOption: "Карамелизация",
        InvalidOptions: []string{"Гомогенизация", "Развитие", "Реакция Майяра"},
    },
    "Температура 180-200° характерна для этого этапа обжарки": {
        CorrectOption: "Первый крек",
        InvalidOptions: []string{"Развитие", "Второй крек", "Карамелизация"},
    },
    "Третий этап обжарки": {
        CorrectOption: "Реакция Майяра",
        InvalidOptions: []string{"Гомогенизация", "Высушивание", "Карамелизация"},
    },
    "Чем больше идёт эта стадия, тем однородней структура кофе": {
        CorrectOption: "Гомогенизация",
        InvalidOptions: []string{"Высушивание", "Карамелизация", "Развитие"},
    },
    "Чем дольше длится данная стадия, тем меньше кислотности и больше горечи будет во вкусе": {
        CorrectOption: "Развитие",
        InvalidOptions: []string{"Карамелизация", "Реакция Майяра", "Первый крек"},
    },
    "Четвёртый этаж обжарки": {
        CorrectOption: "Карамелизация",
        InvalidOptions: []string{"Второй крек", "Развитие", "Первый крек"},
    },
    "Шестой этап обжарки": {
        CorrectOption: "Развитие",
        InvalidOptions: []string{"Второй крек", "Первый крек", "Карамелизация"},
    },

})