import 'dart:convert';
import 'dart:html';

import 'package:dart_toast/dart_toast.dart';
import 'package:mango_entity/bodies/consent.dart';
import 'package:mango_entity/secureapi.dart';

class ConsentForm /*extends FormState*/ {
  UListElement lstConcerns;
  HiddenInputElement hdnCallback;

  ConsentForm(String idElem, String btnSubmit) /*: super(idElem, btnSubmit) */ {
    lstConcerns = querySelector("#lstConcerns");
    hdnCallback = querySelector("#hdnCallback");
    querySelector(btnSubmit).onClick.listen(onSubmitClick);
  }

  Map<String, bool> get concerns {
    final res = new Map<String, bool>();
    res.addEntries(lstConcerns.children.map((e) {
      final label = e.children.first as LabelElement;
      final input = label.children.first as CheckboxInputElement;

      return MapEntry<String, bool>(input.value, input.checked);
    }));
    return res;
  }

  String get usertoken {
    final qs = Uri.splitQueryString(window.location.search);
    return qs["partial"];
  }

  String get state {
    final qs = Uri.splitQueryString(window.location.search);
    return qs["state"];
  }

  String get callback {
    return hdnCallback.value;
  }

  void onSubmitClick(MouseEvent e) async {
    //if (!isFormValid()) {
    //return;
    //}

    //disableSubmit(true);

    final obj = new Consent(usertoken, concerns);

    HttpRequest req = await sendConsent(obj);

    if (req.status == 200) {
      final result = jsonDecode(req.response);
      window.location.href = "${callback}?token=${result}&state=${state}";
    } else {
      new Toast.error(
          title: "Failed!",
          message: req.response,
          position: ToastPos.bottomLeft);
    }
  }
}
