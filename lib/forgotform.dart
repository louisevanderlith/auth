import 'dart:convert';
import 'dart:html';

import 'package:dart_toast/dart_toast.dart';
import 'package:mango_entity/secureapi.dart';
import 'package:mango_ui/formstate.dart';

class ForgotForm extends FormState {
  EmailInputElement _identity;

  ForgotForm(String idElem, String identityElem, String submitBtn)
      : super(idElem, submitBtn) {
    _identity = querySelector(identityElem);

    querySelector(submitBtn).onClick.listen(onSend);
  }

  String get identity {
    return _identity.value;
  }

  void onSend(Event e) async {
    if (isFormValid()) {
      disableSubmit(true);

      var req = await sendForgot(identity);
      var obj = jsonDecode(req.response);

      if (req.status == 200) {
        new Toast.success(
            title: "Success!", message: obj, position: ToastPos.bottomLeft);
        super.form.reset();
      } else {
        new Toast.error(
            title: "Error!",
            message: req.response,
            position: ToastPos.bottomLeft);
      }
    }
  }
}
