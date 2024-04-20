

function sendRequestID() {
    var id = document.getElementById("request-id-input").value;
    var xhr = new XMLHttpRequest();
    xhr.open("GET", "http://localhost:8080/arifmetic", true);
    xhr.setRequestHeader("id", id);
    xhr.onreadystatechange = function () {
      var responsemass = JSON.parse(xhr.response);
      var table = document.getElementById("table-id");
      table.innerHTML = "<tr><th>ID</th><th>Выражение</th><th>Ответ</th><th>Статус</th></tr>";
      for (var i = 0; i < responsemass.length; i++) {
        var response = responsemass[i];
        var newrow = table.insertRow(1);
        newrow.insertCell(0).innerHTML = response.id;
        newrow.insertCell(1).innerHTML = response.task;
        newrow.insertCell(2).innerHTML = response.answer;
        newrow.insertCell(3).innerHTML = response.status;
      };
    };
    xhr.send(null);
  }


  function sendExpressions() {
    var expressions = document.getElementById("expressions").value;
    var expressionList = expressions.split(';').map(expression => expression.trim());
    var data = expressionList;
    var xhr = new XMLHttpRequest();
    xhr.open("POST", "http://localhost:8080/arifmetic", true);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.onreadystatechange = function () {
      var response = xhr.responseText;
      document.getElementById("request-id").textContent = response;
    };
    xhr.send(JSON.stringify(data));
  }
