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
    alertify.error("el campo cantidad no tiene el valor correcto")
  }

  else{

  var texto   = document.getElementById("myTextArea").value;
  var hilos   = document.getElementById("hilos").value;
  var url   = document.getElementById("balanceador").value;
  //console.log(texto);
  if(hilos ==="" || hilos === "0" || url ===""){
    alertify.error("los campos no tiene el valor correcto")
  }
  else {

  for(i=0;i<hilos; i++ ){

    login(cantidad, url, texto);
  }
}
}
  //var tam = $.parseJSON(texto).length
//  console.log(tam);

});

function login(cantidad, url, texto) {

  //console.log();

    data     = '{"Personas":'+texto+',"Parametro":{"Url": "'+url+'","Solicitud": '+cantidad+'}}';

    $.ajax({
        type: 'POST',
        url:  "http://localhost:4000/parametro",
        contentType: "application/json",
        dataType: 'json',
        crossDomain: true,
        async: false,
        data: data,
        success: function (data) {

        //  alert(data);
          if(data === 200){
            alertify.success("Se ejecuto correctamente");
          }



        },
        error: function(data)
        {
         alertify.error("Ocurrio un error");
        }
    });

}
