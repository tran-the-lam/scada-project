import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:intl/intl.dart';


enum SensorType { temperature, humidity, pressure }

class MonitorPage extends StatefulWidget {
  const MonitorPage({super.key, required this.title});
  final String title;

  @override
  State<MonitorPage> createState() => _MonitorList();
}

class _MonitorList extends State<MonitorPage> {
  final List<Map<String, dynamic>> items = [
    {
      'sensor_id': 'S01',
      'parameter': 'Nhiệt độ',
      'value': 50,
      'threshold': 40,
      'created_at': '2023-10-10 10:10:10'
    },
    {
      'sensor_id': 'S02',
      'parameter': 'Độ ẩm',
      'value': 20,
      'threshold': 15,
      'created_at': "2023-10-08 09:10:10"
    },
    {
      'sensor_id': 'S03',
      'parameter': 'Nhiệt độ',
      'value': 50,
      'threshold': 40,
      'created_at': "2023-10-08 08:10:10"
    },
    {
      'sensor_id': 'S04',
      'parameter': 'Nhiệt độ',
      'value': 50,
      'threshold': 40,
      'created_at': "2023-10-08 07:10:10"
    },
    {
      'sensor_id': 'S05',
      'parameter': 'Nhiệt độ',
      'value': 50,
      'threshold': 40,
      'created_at': "2023-10-08 06:10:10"
    },
    {
      'sensor_id': 'S06',
      'parameter': 'Nhiệt độ',
      'value': 50,
      'threshold': 40,
      'created_at': "2023-10-08 05:10:10"
    },
    {
      'sensor_id': 'S07',
      'parameter': 'Nhiệt độ',
      'value': 50,
      'threshold': 40,
      'created_at': "2023-10-08 04:10:10"
    },
    {
      'sensor_id': 'S08',
      'parameter': 'Nhiệt độ',
      'value': 50,
      'threshold': 40,
      'created_at': "2023-10-08 03:10:10"
    },
    {
      'sensor_id': 'S09',
      'parameter': 'Nhiệt độ',
      'value': 50,
      'threshold': 40,
      'created_at': "2023-10-08 02:10:10"
    },
    {
      'sensor_id': 'S10',
      'parameter': 'Nhiệt độ',
      'value': 50,
      'threshold': 40,
      'created_at': "2023-10-08 01:10:10"
    },
  ];

  DateFormat formatter = DateFormat('yyyy/MM/dd HH:mm:ss');
  DateTime selectedStartDate = DateTime.now();
  DateTime selectedEndDate = DateTime.now();
  List<SensorType> selectedSensorTypes = [];

  // enum SensorType { temperature, humidity, light, soil }
  // const _MonitorList({super.key, required this.items});

  @override
  Widget build(BuildContext context) {
    const title = 'Monitor';

    return MaterialApp(
      title: title,
      home: Scaffold(
        appBar: AppBar(
          automaticallyImplyLeading: false,
          title: const Text(title),
          actions: [
            PopupMenuButton(
              color: Colors.deepPurple,
              offset: Offset(0, 50),
              itemBuilder: (BuildContext context) {
                // return <PopupMenuEntry>[
                //   PopupMenuItem(
                //     child: IconButton(
                //       icon: Icon(Icons.search),
                //       onPressed: () {
                //         // Thực hiện hành động khi người dùng nhấn vào icon tìm kiếm
                //         Fluttertoast.showToast(
                //             msg: "your message",
                //             toastLength: Toast.LENGTH_SHORT,
                //             gravity: ToastGravity.BOTTOM, // Also possible "TOP" and "CENTER"
                //             backgroundColor: Color.fromRGBO(152, 54, 244, 1) ,
                //             textColor: Colors.white,
                //         );
                //       },
                //     ),
                //   ),
                // ];
                return [
                  PopupMenuItem(
                    child: Container(
                      padding: EdgeInsets.all(10),
                      child: Column(
                        children: [
                          Text('Giao diện',
                              style: TextStyle(fontWeight: FontWeight.bold)),
                          SizedBox(height: 10),
                          Row(
                            children: [
                              // Text('Ngày bắt đầu: '),
                              Text(formatter.format(selectedStartDate)),
                              IconButton(
                                icon: Icon(Icons.calendar_today),
                                onPressed: () async {
                                 final startDate = await showDatePicker(
                                    context: context,
                                    initialDate: selectedStartDate,
                                    firstDate: DateTime(2000),
                                    lastDate: DateTime(2100),
                                  );
                                  if (startDate != null) {
                                    setState(() {
                                      selectedStartDate = startDate;
                                    });
                                  }
                                   
                                },
                              ),
                            ],
                          ),
                          Row(
                            children: [
                              // Text('Ngày kết thúc: '),
                              Text(formatter.format(selectedEndDate)),
                              IconButton(
                                icon: Icon(Icons.calendar_today),
                                onPressed: () async {
                                  final endDate = await showDatePicker(
                                    context: context,
                                    initialDate: selectedEndDate,
                                    firstDate: DateTime(2000),
                                    lastDate: DateTime(2100),
                                  );
                                  if (endDate != null) {
                                    setState(() {
                                      selectedEndDate = endDate;
                                    });
                                  }
                                },
                              ),
                            ],
                          ),
                          SizedBox(height: 40),
                          Text('Chỉ số:',
                              style: TextStyle(fontWeight: FontWeight.bold)),
                          SizedBox(height: 10),
                          CheckboxListTile(
                            value: true,
                            onChanged: (value) {},
                            title: Text('Nhiệt độ'),
                          ),
                          CheckboxListTile(
                            value: false,
                            onChanged: (value) {},
                            title: Text('Độ ẩm'),
                          ),
                          CheckboxListTile(
                            value: true,
                            onChanged: (value) {},
                            title: Text('Áp suất'),
                          ),
                        ],
                      ),
                    ),
                  ),
                ];
              },
            ),
          ],
        ),
        body: ListView.builder(
          itemCount: items.length,
          itemBuilder: (context, index) {
            return Card(
              child: Column(
                children: [
                  Container(
                    padding: const EdgeInsets.all(8.0),
                    alignment: Alignment.centerLeft,
                    child: Text("Mã cảm biến: " + items[index]['sensor_id']),
                  ),
                  Container(
                    padding: EdgeInsets.all(8.0),
                    child: Text("Chỉ số: " + items[index]['parameter']),
                    alignment: Alignment.centerLeft,
                  ),
                  Container(
                    padding: EdgeInsets.all(8.0),
                    child: Text("Giá trị: " + items[index]['value'].toString()),
                    alignment: Alignment.centerLeft,
                  ),
                  Container(
                    padding: EdgeInsets.all(8.0),
                    child: Text("Ngưỡng: " + items[index]['threshold'].toString()),
                    alignment: Alignment.centerLeft,
                  ),
                  Container(
                    padding: EdgeInsets.all(8.0),
                    child: Text("Thời gian: " + items[index]['created_at']),
                    alignment: Alignment.centerLeft,
                  ),
                ],
              ),
            );
          },
        ),
      ),
    );
  }
}