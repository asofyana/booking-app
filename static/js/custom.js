function ajaxErrorHandler(XMLHttpRequest, textStatus, errorThrown){
    if(XMLHttpRequest.status!=403){
      try{
        var resp = JSON.parse(XMLHttpRequest.responseText);
        if(resp.statuscode==-1){
          window.location.reload();
        }else{
          alert(resp.message);
        }
      }catch(e){
        alert('Sistem sedang mengalami gangguan');
      }
    }
  }
  