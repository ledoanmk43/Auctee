import { useState, useEffect, lazy, Suspense } from 'react';
import { Link as RouterLink, useNavigate, useLocation, useSearchParams, useOutletContext } from 'react-router-dom';
import TimeAgo from 'javascript-time-ago';
import vi from 'javascript-time-ago/locale/vi';
// material
import {
  Card,
  Container,
  Avatar,
  Table,
  TableRow,
  TableBody,
  TableCell,
  Typography,
  TableContainer,
  TablePagination,
  Stack,
} from '@mui/material';

import { styled, useTheme, alpha } from '@mui/material/styles';
import { Icon } from '@iconify/react';
import Scrollbar from '../components/Scrollbar';
import { UserListHead } from '../sections/@dashboard/user';
import { ProductSort, ProductList, ProductFilterSidebar } from '../sections/@dashboard/products';

const Page = lazy(() => import('../components/Page'));
const UpdateProfileForm = lazy(() => import('../sections/update-profile'));

const RootStyle = styled('div')(({ theme }) => ({
  [theme.breakpoints.up('md')]: {
    display: 'flex',
    flexDirection: 'column',
    mx: '176x',
    backgroundColor: 'white',
    height: '100%',
  },
}));

const TABLE_HEAD = [
  { id: 'name', label: 'Tài khoản', alignRight: false },
  { id: 'email', label: 'Email', alignRight: false },
  { id: 'phone', label: 'Số điện thoại', alignRight: false },
  { id: 'time', label: 'Thời gian tạo tài khoản', alignRight: false },
  { id: 'type', label: 'Loại tài khoản', alignRight: false },
];

const EditButton = styled('button')(({ theme }) => ({
  background: `${theme.palette.background.main}`,
  width: 60,
  height: 40,
  border: 'none',
  opacity: 0.9,
  color: 'white',
  fontSize: '1.2rem',
  borderRadius: 0.2,
  ':disabled': {
    opacity: 0.7,
  },
  '&:hover': {
    opacity: 0.8,
  },
}));

TimeAgo.addLocale(vi);
export default function Accounts() {
  const timeAgo = new TimeAgo('vi-VN');
  const userData = useOutletContext();
  const theme = useTheme();
  const navigate = useNavigate();
  const [searchParams, setSearchParams] = useSearchParams();
  const userId = searchParams.get('id');

  // Users table
  const [page, setPage] = useState(0);

  const [order, setOrder] = useState('asc');

  const [selected, setSelected] = useState([]);

  const [orderBy, setOrderBy] = useState('name');

  const [rowsPerPage, setRowsPerPage] = useState(10);
  const handleChangePage = (event, newPage) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (event) => {
    setRowsPerPage(parseInt(event.target.value, 10));
    setPage(0);
  };
  const emptyRows = page > 0 ? Math.max(0, (1 + page) * rowsPerPage - userData.users_list.length) : 0;

  const [auctionsData, setAuctionData] = useState();

  const handleFetchAuctionData = async () => {
    await fetch(`http://localhost:8080/auctee/auctions?page=${1}`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
      mode: 'cors',
    }).then((res) => {
      if (res.status === 200) {
        res.json().then((data) => {
          setAuctionData(data);
        });
      }
      if (res.status === 401) {
        alert('You need to login first');
        navigate('/auctee/login', { replace: true });
      }
    });
  };
  useEffect(() => {
    // eslint-disable-next-line no-unused-expressions
    !auctionsData && handleFetchAuctionData();
  }, [auctionsData]);

  return (
    <Suspense startTransition callback={<></>}>
      <Page sx={{ minHeight: 580 }} title="Accounts Dashboard">
        <RootStyle sx={{ px: 3, py: 2 }}>
          {/* Heading */}
          <Stack direction="row">
            <Avatar sx={{ width: 67, height: 67 }} src={userData?.avatar} alt="photoURL" />
            <Stack sx={{ mx: 2 }} alignItems="flex-start">
              <Typography fontSize={'1.2rem'} variant="body2" sx={{ color: 'black' }}>
                {userData?.lastname} &nbsp;
                {userData?.firstname}
              </Typography>
              <Typography variant="caption" sx={{ color: 'black', position: 'relative' }}>
                Điểm uy tín : {userData?.honor_point} &nbsp;
              </Typography>
              <Typography
                variant="button"
                sx={{
                  textTransform: 'none',
                  bgcolor: '#F62217',
                  color: 'white',
                  borderRadius: 0.5,
                  fontSize: '0.7rem',
                  px: 0.5,
                  mr: 1.5,
                }}
              >
                Quản trị viên
              </Typography>
            </Stack>
            {/* Auctions */}
            <Stack sx={{ ml: 1, borderLeft: '1px solid #f0f0f1' }} alignItems="flex-start">
              {/* Auctions count */}
              <Stack direction="row" alignItems="center">
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ mx: 2, color: 'black' }}>
                  Tồng thu nhập ước tính:
                </Typography>
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ color: theme.palette.background.main }}>
                  {userData?.system_balance.toLocaleString('tr-TR', {
                    style: 'currency',
                    currency: 'VND',
                  })}
                </Typography>
              </Stack>
              {/* Reply rate */}
              <Stack direction="row" alignItems="center">
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ mx: 2, color: 'black' }}>
                  Tổng số lượng người dùng:
                </Typography>
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ color: theme.palette.background.main }}>
                  {userData.total_user} tài khoản
                </Typography>
              </Stack>
              {/* Date Join */}
              <Stack direction="row" alignItems="center">
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ mx: 2, color: 'black' }}>
                  Tham gia:
                </Typography>
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ color: theme.palette.background.main }}>
                  {userData.created_at && timeAgo.format(new Date(userData.created_at))}
                </Typography>
              </Stack>
            </Stack>
            {/* Contact */}
            <Stack sx={{ ml: 4, borderLeft: '1px solid #f0f0f1' }} alignItems="flex-start">
              {/* Auctions count */}
              <Stack direction="row" alignItems="center">
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ mx: 2, color: 'black' }}>
                  Liên hệ:
                </Typography>
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ color: theme.palette.background.main }}>
                  {userData?.phone}
                </Typography>
              </Stack>
              {/* Auctions count */}
              <Stack direction="row" alignItems="center">
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ mx: 2, color: 'black' }}>
                  Email:
                </Typography>
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ color: theme.palette.background.main }}>
                  {userData?.email}
                </Typography>
              </Stack>
              {/* Reply rate */}
              <Stack direction="row" alignItems="center">
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ mx: 2, color: 'black' }}>
                  Địa chỉ:
                </Typography>
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ color: theme.palette.background.main }}>
                  ********, HCM
                </Typography>
              </Stack>
            </Stack>
          </Stack>
        </RootStyle>
        <Container sx={{ my: 3, px: '0 !important' }}>
          <Typography variant="h5" sx={{ mb: 2 }}>
            Danh sách tất cả tài khoản: {userData.total_user}
          </Typography>
          <Container sx={{ px: '0 !important', bgcolor: 'transparent' }}>
            <Card
              sx={{
                // bgcolor: alpha(theme.palette.background.main, 0.05),
                borderRadius: 0,
                // border: `1px solid ${theme.palette.background.main}`,
              }}
            >
              <Scrollbar sx={{ maxHeight: '300px' }}>
                <TableContainer sx={{ minWidth: '100%' }}>
                  <Table>
                    <UserListHead
                      order={order}
                      orderBy={orderBy}
                      headLabel={TABLE_HEAD}
                      rowCount={userData.users_list.length}
                      numSelected={selected.length}
                    />
                    <TableBody>
                      {userData.users_list
                        .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
                        .map((row, index) => {
                          const user = {
                            userId: row.id,
                            shopname: row.shopname || 'unknown_account',
                            email: row.email,
                            phone: row.phone,
                            timeCreated: timeAgo.format(new Date(row.created_at)),
                            role: row.role,
                          };
                          const isItemSelected = selected.indexOf(userId) !== -1;

                          return (
                            <TableRow
                              hover
                              key={index}
                              tabIndex={-1}
                              role="checkbox"
                              selected={isItemSelected}
                              aria-checked={isItemSelected}
                            >
                              <TableCell component="th" scope="row" padding="none" sx={{ pl: 2 }}>
                                <Stack direction="row" alignItems="center" spacing={2}>
                                  {/* <Avatar alt={nickname} src={avatarUrl} /> */}
                                  <Typography
                                    sx={{
                                      color: 'inherit',
                                    }}
                                    variant="subtitle2"
                                    noWrap
                                  >
                                    {user.shopname}
                                  </Typography>
                                  {user.userId === userData.id && (
                                    <Typography variant="body2" sx={{ fontStyle: 'italic' }}>
                                      (Tôi)
                                    </Typography>
                                  )}
                                </Stack>
                              </TableCell>
                              <TableCell align="left">{user.email}</TableCell>
                              <TableCell align="left">{user.phone}</TableCell>
                              <TableCell align="left">{user.phone}</TableCell>
                              <TableCell align="left">{user.role === 1 ? 'Quản trị viên' : 'Người dùng'}</TableCell>
                            </TableRow>
                          );
                        })}
                      {emptyRows > 0 && (
                        <TableRow style={{ height: 53 * emptyRows }}>
                          <TableCell colSpan={6} />
                        </TableRow>
                      )}
                    </TableBody>
                  </Table>
                </TableContainer>
              </Scrollbar>

              <TablePagination
                rowsPerPageOptions={[4, 8, 24]}
                component="div"
                count={userData.users_list.length}
                rowsPerPage={rowsPerPage}
                page={page}
                onPageChange={handleChangePage}
                onRowsPerPageChange={handleChangeRowsPerPage}
              />
            </Card>
          </Container>
        </Container>
      </Page>
    </Suspense>
  );
}
