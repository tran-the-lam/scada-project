import 'package:http/http.dart' as http;
import 'package:mydemo/utils/constant.dart';
import 'dart:convert';

Future<String> login(String userID, String pwd, String deviceInfo) async {
  var url = Uri.parse("${Constant.BASE_URL}/login");
  print(url);
  var body = json.encode({"user_id": userID, "password": pwd, "device_id": Constant.DEVICE_ID});
  var headers = {
    "Content-Type": "application/json",
    "User-Agent": deviceInfo,
  };
  
  var response = await http.post(url, body: body, headers: headers);
  
  if (response.statusCode == 200) {
    // Xử lý dữ liệu trong response.body
    print(response.body);
    final token = json.decode(response.body)['token'];
    return token;
  }
  
  return "";
}

Future<void> sendDataToApi() async {
  var url = Uri.parse("https://example.com/api/data");
  
  // Tạo body dữ liệu
  var body = {"name": "John", "age": 25};
  
  // Gửi request POST với body dữ liệu
  var response = await http.post(url, body: body);
  
  if (response.statusCode == 200) {
    // Xử lý dữ liệu trong response.body
    print(response.body);
  } else {
    // Xử lý lỗi
    print('Lỗi ${response.statusCode}');
  }
}