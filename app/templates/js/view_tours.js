function createTourInfo(tourFacts) {
    let tourInfo = document.createElement('div');
    tourInfo.classList.add('tour-info');

    let factList = document.createElement('ul');
    for (var key in tourFacts ) {
        let tourFact = document.createElement('li');
        tourFact.textContent = key + ': ' + tourFacts[key];
        factList.appendChild(tourFact);
    }

    tourInfo.appendChild(factList)

    return tourInfo
}

function createTourPrice(tourInfo) {
    let tourPrice = document.createElement('div');
    tourPrice.classList.add('tour-price');

    for (var key in tourInfo ) {
        let tourFact = document.createElement('p');
        tourFact.textContent = key + ": " + tourInfo[key];
        tourPrice.appendChild(tourFact);
    }

    return tourPrice
}

function createTourDetails(from, date, days, type, cost, tourists) {
    let tourInfo = {
        "Место отправления": from,
        "Дата отправления": date,
        "Продолжительность": days,
        "Тип отдыха": type,
    };

    let tourPrice = {
        "Стоимость": cost,
        "Количество туристов": tourists,
    };

    let tourInfoDetails = createTourInfo(tourInfo);
    let tourPriceDetails = createTourPrice(tourPrice);

    let tourDetails = document.createElement('div');
    tourDetails.classList.add('tour-datails');
    tourDetails.appendChild(tourInfoDetails);
    tourDetails.appendChild(tourPriceDetails);

    return tourDetails;
}


function createTourCard(from, to, cost, date, days, tourists, type) {
    let tourCard = document.createElement('div');
    tourCard.classList.add('tour-card');

    let tourName = document.createElement('h3');
    tourName.textContent = to;

    tourCard.appendChild(tourName);

    let tourDetails = createTourDetails(from, date, days, type, cost, tourists);

    let buttonAbout = document.createElement('a');
    buttonAbout.classList.add('tour-button');
    buttonAbout.href = '#';
    buttonAbout.textContent = "Подробнее";

    tourCard.appendChild(tourDetails);
    tourCard.appendChild(buttonAbout);

    return tourCard
}

