
# Xây dựng hệ thống quản trị Reward(Offer) cơ bản

## Bài toán
Xây dựng hệ thống quản trị Quà tặng (Program) cho khách hàng ViettelPay. Khi khách hàng ViettelPay thực hiện các loại giao dịch nhất định thì tặng cho khách hàng các quà tặng (Reward) tương ứng (Định nghĩa cụ thể ở phần sau). 

## Đối tượng của hệ thống
### End user
Là người dùng cuối sử dụng dịch vụ ViettelPay và là đối tượng được hưởng các chính sách quà tặng khi phát sinh giao dịch. 
### Partner 
Là các đối tác cung cấp quà tặng của ViettelPay, đối tác có thể khai báo chương trình quà tặng của End user. Chương trình quà tặng chỉ có hiệu lực sau khi Partner khai báo và được VDS Staff phê duyệt. Mỗi đối tác chị 
### VDS Staff
Là nhân viên của VDS có trách nhiệm review và phê duyệt các chương trình quà tặng của Partner. 

## Yêu cầu 
### Story 1 
`Chương trình quà tặng (Program) được kích hoạt thành công khi Partner khai báo chương trình quà tặng cho End User trên web admin và sau đó được phê duyệt thành công bởi VDS Staff`

1. Thông tin của 1 chương trình quà tặng được định nghĩa bao gồm: 
- Tên chương trình (`program_name`)
- Mã đối tác (`partner_code`): Là mã của đối tác đang cung cấp chương trình quà tặng
- Thời gian bắt đầu, kết thúc của chương trình quà tặng (`program_effective_from`, `program_effective_to`)
- Điều kiện của chương trình quà tặng (Policy) có thể bao gồm các loại điều kiện sau
	- Nhân khẩu học: Tuổi/Giới tính (`age`/`gender`)
	- Loại giao dịch: Chọn 1 hoặc nhiều loại giao dịch. 
- Số quà tặng tối đa trong chương trình cho 1 khách hàng (`max_reward_per_user`)
- Danh sách quà tặng(`reward_list`): Là Voucher của Partner, bao gồm thông tin:
	- `reward_code`: Mã voucher 
	- `reward_effective_from`: Thời điểm voucher bắt đầu có hiệu lực (epoch second)
 	- `reward_effective_to`: Thời điểm voucher hết hiệu lực (epoch second)

Ví dụ: CGV muốn gửi Voucher Tiki cho các User nữ nhân ngày 8/3 với khách hàng thực hiện giao dịch Topup di động.

2. VDS Staff có thể nhìn thấy các chương trình quà tặng chờ phê duyệt và có thể chọn ***Phê duyệt*** hoặc ***Từ chối phê duyệt*** Chương trình quà tặng. Chương trình quà tặng có hiệu lực ngay lập tức khi VDS Staff chọn ***Phê duyệt***

### Story 2 
`Khi người dùng thực hiện giao dịch trong 1 khoảng thời gian cụ thể (VD 8/3), nếu thông tin giao dịch và người dùng đáp ứng Điều kiện của chương trình quà tặng, thực hiện tặng quà tặng cho người dùng.`

1. ***Thông tin giao dịch*** được giả định đẩy qua hệ thống Kafka , hệ thống sẽ process các event tương ứng để tính toán các Reward phù hợp.
2. ***Thông tin giao dịch*** bao gồm: 
- `event_name`: Loại giao dịch, gồm các loại sau: 
	- `airtime`:  Nạp tiền topup di động 
	- `loan`:  Thanh toán khoản vay
	- `transfer`:  Chuyển tiền theo số tài khoản
- `properties`: các thông tin của giao dịch, bao gồm: 
	- `event_id`:  của giao dịch
	- `timestamp`: Thời điểm giao dịch được thực hiện thành công, theo định dạng Epoch second. 
- `user_id`: user id của End user. 
- Khi các End user đủ điều kiện nhận Quà tặng, hệ thống gửi email cho khách hàng thông tin voucher. (Giả định đẩy ra CDR log với nội dung):
	- `timestamp`|`event_id` | `reward_code`

## Dữ liệu test 
### Dữ liệu profile của End user
todo: Tạo file dữ liệu người dùng
### Dữ liệu event giả lập qua kafka 
todo: Dựng kafka và push message vào topic trên môi trường test
### Danh sách chương trình quà tặng
todo: Tạo file dữ liệu test
 
## Hướng dẫn nộp bài 
1. Fork repo sau về: https://github.com/weburnit/ddd-java
2. Tạo nhánh tương ứng: <email>
3. Thực hiện code
4. Tạo PR vào nhánh master của repo chính. 

## Tiêu chí đánh giá (bổ sung thêm)
* Các component thể hiện Hexagonal: https://techmaster.vn/posts/34239/kien-truc-luc-giac-trong-xay-dung-ung-dung
* SPI/API/TDD
* ChainOfRes: FraudPolicy & FrequencyCapPolicy
* Design patterns
