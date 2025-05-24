Parking Suggestion Map System
 Giới thiệu
Ứng dụng này hỗ trợ người dùng tìm đường đi và gợi ý chỗ đỗ xe gần nhất dựa trên vị trí hiện tại. Sử dụng thuật toán Dijkstra cho tìm đường và hiển thị bản đồ tương tác bằng LeafletJS.

 Tính năng chính
Tìm đường giữa các bãi đỗ xe có sẵn.

Gợi ý bãi đỗ xe gần nhất còn chỗ trống từ vị trí hiện tại.

Giao diện bản đồ trực quan, có thể nhập vị trí hiện tại thủ công.

Quản lý số lượng xe đang chiếm trong từng nhà xe qua API.

 Cấu trúc hệ thống
Backend: Go (Gin Framework)

Frontend: HTML + JavaScript (LeafletJS)

API: RESTful

 Cài đặt và chạy
1. Clone project:
bash
Sao chép
Chỉnh sửa
git clone https://github.com/your-repo/parking-map.git
cd parking-map
2. Build file thực thi:
bash
Sao chép
Chỉnh sửa
go build -o parking.exe
3. Chạy chương trình:
bash
Sao chép
Chỉnh sửa
./parking.exe
Ứng dụng sẽ chạy tại http://localhost:8080.

API
1. Cập nhật số lượng xe đã chiếm:
POST /update-occupied

json
Sao chép
Chỉnh sửa
{
  "id": "Nha xe D9",
  "occupied": 150
}
2. Tìm đường:
GET /path?from=Nha xe D9&to=Nha xe B6

3. Gợi ý bãi đỗ xe gần nhất:
POST /nearest

json
Sao chép
Chỉnh sửa
{
  "lat": 21.0045,
  "lng": 105.8456
}
 Triển khai tĩnh
Đảm bảo file index.html nằm cùng thư mục với file .exe, vì ứng dụng sử dụng router.StaticFile("/", "./index.html") để phục vụ giao diện người dùng.



 Tham khảo
LeafletJS

Gin Gonic

