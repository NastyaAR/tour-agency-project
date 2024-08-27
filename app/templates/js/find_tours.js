function getIntFromString(str) {
    const result = parseInt(str, 10);
    return isNaN(result) ? -1 : result;
}

function create_tour_criteria_json(event) {
    let jsonData = {};

    console.log(document.getElementById('chillPlace').value)
    jsonData['chillPlace'] = document.getElementById('chillPlace').value;
    jsonData['fromPlace'] = document.getElementById('fromPlace').value;
    jsonData['date'] = document.getElementById('date').value;
    if (jsonData['date'] == "") {
        jsonData['date'] = "1970-01-01T00:00"
    }

    jsonData['duration'] = document.getElementById('duration').value;
    jsonData['duration'] = getIntFromString(jsonData['duration'])

    jsonData['cost'] = document.getElementById('cost').value;
    jsonData['cost'] = getIntFromString(jsonData['cost'])

    jsonData['touristsNumber'] = document.getElementById('touristsNumber').value;
    jsonData['touristsNumber'] = getIntFromString(jsonData['touristsNumber'])

    jsonData['chillType'] = document.getElementById('chillType').value;

    let jsonString = JSON.stringify(jsonData);
    console.log(jsonData)
    console.log(jsonString)
    return jsonString;
}

function tour_submitter(event) {
    event.preventDefault(); // Предотвращаем стандартное поведение формы

    fetch('http://0.0.0.0:8080/find/tours', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: create_tour_criteria_json(event),
    })
        .then(response => {
            if (response.redirected) {
                window.location.href = response.url; // Перенаправляем на новый URL, если сервер отправляет редирект
            } else {
                return response.json();
            }
        })
        .then(data => {
            console.log('Успех:', data);
            // Здесь можно обработать данные, если необходимо
        })
        .catch((error) => {
            console.error('Ошибка:', error);
        });
}
