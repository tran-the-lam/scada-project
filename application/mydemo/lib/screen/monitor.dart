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
  final List<http.Event> items = [];

  DateTime _selectedStartDate = DateTime.now();
  DateTime _selectedEndDate = DateTime.now();

  List<String> options = ['temperature', 'humidity'];
  String _parameter = 'temperature';

  TextEditingController _sensorIDController = TextEditingController();

  @override
  void initState() {
    super.initState();
    fetchEvent();
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

  Future<void> searchEventOnPress() async {
    if (_sensorIDController.text.isEmpty) {
      Fluttertoast.showToast(
        msg: "Vui lòng nhập mã cảm biến",
        toastLength: Toast.LENGTH_SHORT,
        gravity: ToastGravity.CENTER,
      );
      return;
    }
    final events = await http.searchEvent(_sensorIDController.text);
    setState(() {
      items.clear();
      items.addAll(events);
    });

    Navigator.of(context).pop();
  }

  Future<void> searchCancelBtnOnPress() async {
    final events = await http.getAllEvent();
    setState(() {
      items.clear();
      items.addAll(events);
    });

    Navigator.of(context).pop();
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
                    
                    // SizedBox(height: 10, width: 10),
                    // Row(
                    //     children: [
                    //       Text('Chỉ số:', style: TextStyle(fontWeight: FontWeight.bold)),
                    //       SizedBox(height: 10, width: 10),
                    //       DropdownButton(
                    //         value: _parameter,
                    //         items: options.map((String option) {
                    //           return DropdownMenuItem(
                    //             value: option,
                    //             child: Text(option),
                    //           );
                    //         }).toList(),
                    //         onChanged: (String? newValue) {
                    //           setState(() {
                    //             // print(newValue);
                    //             _parameter = newValue!;
                    //             // print('Role $_parameter');
                    //           });
                    //         },
                    //       ),
                    //   ]
                    // ),
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
                  onPressed: searchEventOnPress,
                  child: Text('Tìm kiếm'),
                ),
                ElevatedButton(
                  onPressed: searchCancelBtnOnPress,
                  child: Text('Hủy'),
                ),
              ],
            );
          },
        );
      },
    );
  }

  void fetchEvent() async {
    final events = await http.getAllEvent();
    setState(() {
      items.clear();
      items.addAll(events);
    });
  }

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
          child: ListView.builder(
                    itemCount: items.length,
                    itemBuilder: (context, index) {
                      final item = items[index];
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
                  ), 
                // }
              // },
        ),
      ),
      );
  }
}