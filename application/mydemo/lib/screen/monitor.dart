import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:intl/intl.dart';
import 'package:flutter_datetime_picker/flutter_datetime_picker.dart' as fdp;
import 'package:mydemo/utils/http.dart' as http;

class MonitorPage extends StatefulWidget {
  const MonitorPage({super.key, required this.title});
  final String title;

  @override
  State<MonitorPage> createState() => _MonitorList();
}

class _MonitorList extends State<MonitorPage> {
  final List<http.Event> items = [
    // {
    //   'sensor_id': 'S01',
    //   'parameter': 'Nhiệt độ',
    //   'value': 50,
    //   'threshold': 40,
    //   'created_at': '2023-10-10 10:10:10'
    // },
    // {
    //   'sensor_id': 'S02',
    //   'parameter': 'Độ ẩm',
    //   'value': 20,
    //   'threshold': 15,
    //   'created_at': "2023-10-08 09:10:10"
    // },
    // {
    //   'sensor_id': 'S03',
    //   'parameter': 'Nhiệt độ',
    //   'value': 50,
    //   'threshold': 40,
    //   'created_at': "2023-10-08 08:10:10"
    // },
    // {
    //   'sensor_id': 'S04',
    //   'parameter': 'Nhiệt độ',
    //   'value': 50,
    //   'threshold': 40,
    //   'created_at': "2023-10-08 07:10:10"
    // },
    // {
    //   'sensor_id': 'S05',
    //   'parameter': 'Nhiệt độ',
    //   'value': 50,
    //   'threshold': 40,
    //   'created_at': "2023-10-08 06:10:10"
    // },
    // {
    //   'sensor_id': 'S06',
    //   'parameter': 'Nhiệt độ',
    //   'value': 50,
    //   'threshold': 40,
    //   'created_at': "2023-10-08 05:10:10"
    // },
    // {
    //   'sensor_id': 'S07',
    //   'parameter': 'Nhiệt độ',
    //   'value': 50,
    //   'threshold': 40,
    //   'created_at': "2023-10-08 04:10:10"
    // },
    // {
    //   'sensor_id': 'S08',
    //   'parameter': 'Nhiệt độ',
    //   'value': 50,
    //   'threshold': 40,
    //   'created_at': "2023-10-08 03:10:10"
    // },
    // {
    //   'sensor_id': 'S09',
    //   'parameter': 'Nhiệt độ',
    //   'value': 50,
    //   'threshold': 40,
    //   'created_at': "2023-10-08 02:10:10"
    // },
    // {
    //   'sensor_id': 'S10',
    //   'parameter': 'Nhiệt độ',
    //   'value': 50,
    //   'threshold': 40,
    //   'created_at': "2023-10-08 01:10:10"
    // },
  ];

  DateTime _selectedStartDate = DateTime.now();
  DateTime _selectedEndDate = DateTime.now();

  List<String> options = ['Nhiệt độ', 'Độ ẩm'];
  String _parameter = 'Nhiệt độ';

  TextEditingController _sensorIDController = TextEditingController();

  @override
  void initState() {
    super.initState();
    // fetchEvent();
  }
  
  startDateTimePickerWidget(BuildContext context, Function setState) {
    return fdp.DatePicker.showDateTimePicker(
      context,
      minTime: DateTime.now().subtract(Duration(days: 365)),
      maxTime: DateTime.now(),
      currentTime: _selectedStartDate,
      locale: fdp.LocaleType.vi,
      onConfirm: (DateTime dateTime) {
        setState(() {
          _selectedStartDate = dateTime;
        });
        print(_selectedStartDate);
      },
    );
  }

  endDateTimePickerWidget(BuildContext context, Function setState) {
    return fdp.DatePicker.showDateTimePicker(
      context,
      minTime: DateTime.now().subtract(Duration(days: 365)),
      maxTime: DateTime.now(),
      currentTime: _selectedEndDate,
      locale: fdp.LocaleType.vi,
      onConfirm: (DateTime dateTime) {
        setState(() {
          _selectedEndDate = dateTime;
        });
        print(_selectedEndDate);
      },
    );
  }

  void showSearchPopup(BuildContext context) {
    showDialog(
      context: context,
      builder: (BuildContext context) {
        return StatefulBuilder(builder: (context, setState) {
            return AlertDialog(
              title: Text('Tìm kiếm dữ liệu', style: TextStyle(fontWeight: FontWeight.bold), textAlign: TextAlign.center),
              content: Column(
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    // Row(
                    //   children: [
                    //     Text('Ngày bắt đầu: ${DateFormat('dd-MM-yyyy HH:mm').format(_selectedStartDate)}', style: TextStyle(fontWeight: FontWeight.bold)),
                    //     IconButton(
                    //     icon: Icon(Icons.calendar_today),
                    //     onPressed: () {
                    //       startDateTimePickerWidget(context, setState);
                    //     },
                    //   ),
                    //   ]
                    // ),
                    // Row(
                    //   children: [
                    //     Text('Ngày kết thúc: ${DateFormat('dd-MM-yyyy HH:mm').format(_selectedEndDate)}', style: TextStyle(fontWeight: FontWeight.bold)),
                    //     IconButton(
                    //     icon: Icon(Icons.calendar_today),
                    //     onPressed: () {
                    //       endDateTimePickerWidget(context, setState);
                    //     },
                    //   ),
                    //   ]
                    // ),
                    
                    SizedBox(height: 10, width: 10),
                    Row(
                        children: [
                          Text('Chỉ số:', style: TextStyle(fontWeight: FontWeight.bold)),
                          SizedBox(height: 10, width: 10),
                          DropdownButton(
                            value: _parameter,
                            items: options.map((String option) {
                              return DropdownMenuItem(
                                value: option,
                                child: Text(option),
                              );
                            }).toList(),
                            onChanged: (String? newValue) {
                              setState(() {
                                // print(newValue);
                                _parameter = newValue!;
                                // print('Role $_parameter');
                              });
                            },
                          ),
                      ]
                    ),
                    SizedBox(height: 10, width: 10),
                    Row(
                      children: [
                        Expanded(
                          child: TextField(
                            controller: _sensorIDController,
                            decoration: InputDecoration(hintText: "Nhập mã cảm biến"),
                          ),
                        ),
                      ]
                    ),
                ],
              ), 
              actions: [
                ElevatedButton(
                  onPressed: () {
                    // Thực hiện hành động tìm kiếm ở đây
                    // Ví dụ: gọi một API để tìm kiếm theo từ khóa
                    // print('Mã nhân viên: $userID');
                    Navigator.of(context).pop();
                  },
                  child: Text('Tìm kiếm'),
                ),
                ElevatedButton(
                  onPressed: () {
                    // Đóng popup khi nhấn nút Hủy
                    Navigator.of(context).pop();
                  },
                  child: Text('Hủy'),
                ),
              ],
            );
          },
        );
      },
    );
  }

  // void fetchEvent() async {
  //   final events = await http.getAllEvent();
  //   setState(() {
  //     items.clear();
  //     items.addAll(events);
  //   });
  // }
  @override
  Widget build(BuildContext context) {
    const title = 'Monitor';
    List<String> options = ['Nhiệt độ', 'Độ ẩm'];
    

    return MaterialApp(
      title: title,
      debugShowCheckedModeBanner: false,
      home: Scaffold(
        appBar: AppBar(
          automaticallyImplyLeading: false,
          title: const Text(title),
          actions: [
            IconButton(
              icon: Icon(Icons.search),
              onPressed: () {
                showSearchPopup(context);
              },
            ),
          ],
        ),
        body: Container(
          padding: EdgeInsets.all(16.0),
          child: FutureBuilder(
            future: http.getAllEvent(),
            builder: (BuildContext ctx, AsyncSnapshot snapshot) {
                if (snapshot.connectionState == ConnectionState.waiting) {
                  return Center(
                    child: CircularProgressIndicator(),
                  );
                } else if (snapshot.hasError) {
                  return Center(
                    child: Text('Error: ${snapshot.error}'),
                  );
                } else if (!snapshot.hasData) {
                  return Center(
                    child: Text('No data available'),
                  );
                } else {
                  return ListView.builder(
                    itemBuilder: (context, index) {
                      final item = snapshot.data[index];
                      // print(snapshot.data[index]);
                      return Card(
                        child: Column(
                          children: [
                            Container(
                              padding: const EdgeInsets.all(8.0),
                              alignment: Alignment.centerLeft,
                              child: Text("Mã cảm biến: " + item.sensorId),
                            ),
                            Container(
                              padding: EdgeInsets.all(8.0),
                              child: Text("Chỉ số: " + item.parameter),
                              alignment: Alignment.centerLeft,
                            ),
                            Container(
                              padding: EdgeInsets.all(8.0),
                              child: Text("Giá trị: " + item.value.toString()),
                              alignment: Alignment.centerLeft,
                            ),
                            Container(
                              padding: EdgeInsets.all(8.0),
                              child: Text("Ngưỡng: " + item.threshold.toString()),
                              alignment: Alignment.centerLeft,
                            ),
                            Container(
                              padding: EdgeInsets.all(8.0),
                              child: Text("Thời gian: " + item.createdAt),
                              alignment: Alignment.centerLeft,
                            ),
                          ],
                        ),
                      );
                  },
                  ); 
                }
              },
        ),
      ),
      ),
    );
  }
}