import 'package:http/http.dart' as http;
import 'package:mydemo/utils/constant.dart';
import 'package:mydemo/utils/utils.dart' as utils;

import 'dart:convert';

Future<String> login(String userID, String pwd, String deviceInfo) async {
  var url = Uri.parse("${Constant.BASE_URL}/login");
  var body = json.encode(
      {"user_id": userID, "password": pwd, "device_id": Constant.DEVICE_ID});
  var headers = {
    "Content-Type": "application/json",
    "User-Agent": deviceInfo,
    "Remote-Address": ""
  };

  var response = await http.post(url, body: body, headers: headers).timeout(
    const Duration(seconds: 5),
    onTimeout: () {
      // Time has run out, do what you wanted to do.
      return http.Response('Error', 408); // Request Timeout response status code
    },
  );

  if (response.statusCode == 200) {
    // Xử lý dữ liệu trong response.body
    print(response.body);
    final token = json.decode(response.body)['token'];
    await utils.saveUserInfo(userID, token);
    return token;
  }

  return "";
}

class Event {
  final String sensorId;
  final String parameter;
  final int value;
  final int threshold;
  final String createdAt;

  Event({
    required this.sensorId,
    required this.parameter,
    required this.value,
    required this.threshold,
    required this.createdAt,
  });
}

Future<List<Event>> getAllEvent() async {
  String url = "${Constant.BASE_URL}/events";
  final token = await utils.getToken();
  Map<String, String> requestHeaders = {
    'Authorization': 'Bearer $token',
    'Accept': '*/*',
    "Access-Control-Allow-Origin": "*",
  };

  print("url: $url \nheaders: $requestHeaders");

  
  final response = await http.get(Uri.parse(url), headers: requestHeaders);
  var responseData = json.decode(response.body)["data"];

  List<Event> events = [];
  for (var item in responseData) {
    final dateTime =
        DateTime.fromMillisecondsSinceEpoch(item['timestamp'], isUtc: true);
    Event event = Event(
      sensorId: item['sensor_id'],
      parameter: item['parameter'],
      value: item['value'],
      threshold: item['threshold'],
      createdAt: dateTime.toIso8601String(),
    );
    events.add(event);
  }

  return events;
}

Future<List<Event>> searchEvent(String sensorID) async {
  String url = "${Constant.BASE_URL}/events/search?sensor_id=$sensorID";
  final token = await utils.getToken();
  Map<String, String> requestHeaders = {
    'Authorization': 'Bearer $token',
    'Accept': '*/*',
    "Access-Control-Allow-Origin": "*",
  };

  final response = await http.get(Uri.parse(url), headers: requestHeaders);
  var responseData = json.decode(response.body)["data"];

  List<Event> events = [];
  for (var item in responseData) {
    final dateTime =
        DateTime.fromMillisecondsSinceEpoch(item['timestamp'], isUtc: true);
    Event event = Event(
      sensorId: item['sensor_id'],
      parameter: item['parameter'],
      value: item['value'],
      threshold: item['threshold'],
      createdAt: dateTime.toIso8601String(),
    );
    events.add(event);
  }

  return events;
}


class User {
  final String userID;
  final String role;
  final int status;

  User({
    required this.userID,
    required this.role,
    required this.status,
  });
}

Future<List<User>> getAllUser() async {
  String url = "${Constant.BASE_URL}/users";
  final token = await utils.getToken();
  Map<String, String> requestHeaders = {
    'Authorization': 'Bearer $token',
    'Accept': '*/*',
    "Access-Control-Allow-Origin": "*",
  };

  final response = await http.get(Uri.parse(url), headers: requestHeaders);
  var responseData = json.decode(response.body)["data"];

  List<User> users = [];
  for (var item in responseData) {
    User u = User(
      userID: item['user_id'],
      role: item['role'],
      status: item['status'],
    );
    users.add(u);
  }

  return users;
} 

Future<String> addUser(String userID, String role) async {
  String url = "${Constant.BASE_URL}/users";
  final token = await utils.getToken();
  Map<String, String> requestHeaders = {
    'Authorization': 'Bearer $token',
    'Accept': '*/*',
    "Access-Control-Allow-Origin": "*",
    "Content-Type": "application/json",
  };

  print("Add user url: $url \nheaders: $requestHeaders");
  var body = json.encode({"user_id": userID, "role": role});

  final response = await http.post(Uri.parse(url), headers: requestHeaders, body: body);
  var responseData = json.decode(response.body)["message"];
  return responseData;
}

// Future<void> sendDataToApi() async {
//   var url = Uri.parse("https://example.com/api/data");

//   // Tạo body dữ liệu
//   var body = {"name": "John", "age": 25};

//   // Gửi request POST với body dữ liệu
//   var response = await http.post(url, body: body);

//   if (response.statusCode == 200) {
//     // Xử lý dữ liệu trong response.body
//     print(response.body);
//   } else {
//     // Xử lý lỗi
//     print('Lỗi ${response.statusCode}');
//   }
// }

class HistoryLogin {
  // final String ipAddress;
  final String device;
  final String createdAt;

  HistoryLogin({
    required this.device,
    required this.createdAt,
  });
}

Future<List<HistoryLogin>> GetHistoryLogin() async {
  String url = "${Constant.BASE_URL}/users/history/login";
  final token = await utils.getToken();
  Map<String, String> requestHeaders = {
    'Authorization': 'Bearer $token',
    'Accept': '*/*',
    "Access-Control-Allow-Origin": "*",
  };

  final response = await http.get(Uri.parse(url), headers: requestHeaders);
  var responseData = json.decode(response.body)["data"];

  List<HistoryLogin> hl = [];
  for (var item in responseData) {
    HistoryLogin h = HistoryLogin(
      device: item['user_agent'],
      createdAt: item['time'],
    );
    hl.add(h);
  }

  return hl;
} 


Future<bool> updatePassword(String oldPwd, String newPwd) async {
  var url = Uri.parse("${Constant.BASE_URL}/users/password");
  var body = json.encode({"old_password": oldPwd, "new_password": newPwd});
  var headers = {
    "Content-Type": "application/json",
    "Authorization": "Bearer ${await utils.getToken()}",
  };

  var response = await http.put(url, body: body, headers: headers).timeout(
    const Duration(seconds: 5),
    onTimeout: () {
      // Time has run out, do what you wanted to do.
      return http.Response('Error', 408); // Request Timeout response status code
    },
  );

  print(response.body);
  if (response.statusCode == 200) {
    // Xử lý dữ liệu trong response.body
    
    return true;
  }

  return false;
}


Future<bool> resetPassword(String userId) async {
  var url = Uri.parse("${Constant.BASE_URL}/users/${userId}/reset-password");
  var headers = {
    "Content-Type": "application/json",
    "Authorization": "Bearer ${await utils.getToken()}",
  };

  var response = await http.post(url, headers: headers).timeout(
    const Duration(seconds: 5),
    onTimeout: () {
      // Time has run out, do what you wanted to do.
      return http.Response('Error', 408); // Request Timeout response status code
    },
  );

  print(response.body);
  if (response.statusCode == 200) {
    return true;
  }

  return false;

}