// component
import Iconify from '../../components/Iconify';

// ----------------------------------------------------------------------

const getIcon = (name) => <Iconify icon={name} width={25} height={25} />;

const navConfig = [
  {
    title: 'Tài khoản của tôi',
    icon: getIcon('ant-design:user-outlined'),
    children: [
      { title: 'Hồ sơ', path: '/auctee/user/profile' },
      { title: 'Địa chỉ', path: '/auctee/user/address' },
      { title: 'Đổi mật khẩu', path: '/auctee/user/change-password' },
      { title: 'Thanh toán', path: '/auctee/user/purchase' },
    ],
  },
  {
    title: 'Phiên đấu giá',
    icon: getIcon('ri:auction-line'),
    children: [
      { title: 'Kho', path: '/auctee/user/inventory' },
      { title: 'Thống kê', path: '/auctee/user/dashboard' },
      { title: 'Đơn thanh toán', path: '/auctee/user/order' },
    ],
  },
];

export default navConfig;
