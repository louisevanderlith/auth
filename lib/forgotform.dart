import 'dart:html';

import 'package:Auth.APP/formstate.dart';

class ForgotForm extends FormState{
  EmailInputElement _identity;

  ForgotForm(String idElem, String identityElem, String submitBtn)
  :super(idElem, submitBtn){
    _identity = querySelector(identityElem);

    querySelector(submitBtn).onClick.listen(onSend);
  }
  
  String get identity {
    return _identity.value;
  }

  void onSend(Event e) async {
    if (isFormValid()) {
      disableSubmit(true);

      var result = await sendForgot(identity);
      var obj = jsonDecode(result.response);

      if (result.status == 200) {
        afterSend(obj['Data']);
      } else {
        _error.text = obj['Error'];
      }
    }
  }
  
  void afterSend(String sessionID) {
    var finalURL = window.localStorage['return'];
    finalURL += "?access_token=" + sessionID;

    window.location.replace(finalURL);
  }
}