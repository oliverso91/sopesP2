$(document).ready(function () {





});

document.getElementById('inputfile')
            .addEventListener('change', function() {

            var fr=new FileReader();
            fr.onload=function(){
                document.getElementById('myTextArea')
                        .textContent=fr.result;
            }

            fr.readAsText(this.files[0]);
        })


$("#aceptar").on('click', function () {


  var cantidad =  document.getElementById("numero").value;

  if(cantidad ==="" || cantidad === "0"){
    alertify.error("el campo cantidad esta vacio")
  }

  else{

  var texto   = document.getElementById("myTextArea").value;
  var hilos   = document.getElementById("hilos").value;
  var url   = document.getElementById("balanceador").value;

  //console.log(texto);
    login(cantidad, hilos, url, texto);
}
  //var tam = $.parseJSON(texto).length
//  console.log(tam);

});

function login(cantidad, hilos, url, texto) {

  //console.log();

    data     = '{"Personas":'+texto+',"Parametro":{"Url": "'+url+'","Hilos" : '+hilos+',"Solicitud": '+cantidad+'}}';

    $.ajax({
        type: 'POST',
        url:  "http://localhost:4000/parametro",
        contentType: "application/json",
        dataType: 'json',
        crossDomain: true,
        async: false,
        data: data,
        success: function (data) {

          alert(data);



        },
        error: function(data)
        {
         console.log(data);
        }
    });

}
