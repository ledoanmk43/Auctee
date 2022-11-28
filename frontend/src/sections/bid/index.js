import { useState, useEffect, useCallback } from 'react';
import { Link as RouterLink, useNavigate, useSearchParams, useLocation, useOutletContext } from 'react-router-dom';
import useWebSocket, { ReadyState } from 'react-use-websocket';
import TimeAgo from 'javascript-time-ago';
// English.
import vi from 'javascript-time-ago/locale/vi';
import { Icon } from '@iconify/react';
import {
  Box,
  Card,
  Table,
  Stack,
  TableRow,
  TableBody,
  TableCell,
  Container,
  Typography,
  TableContainer,
  TablePagination,
  Button,
  Dialog,
  useMediaQuery,
} from '@mui/material';
import { styled, useTheme, alpha } from '@mui/material/styles';
import 'moment/locale/vi';
import Scrollbar from '../../components/Scrollbar';
import { UserListHead } from '../@dashboard/user';

const TABLE_HEAD = [
  { id: 'name', label: 'Người tham gia', alignRight: false },
  { id: 'value', label: 'Số tiền ', alignRight: false },
  { id: 'time', label: 'Thời gian ', alignRight: false },
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

const StyledButton = styled('button')(({ theme }) => ({
  background: `${theme.palette.background.main}`,
  display: 'flex',
  justifyContent: 'center',
  alignItems: 'center',
  width: 160,
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
    opacity: 0.85,
  },
}));

TimeAgo.addLocale(vi);
const BidSection = ({ product, auction }) => {
  const timeAgo = new TimeAgo('vi-VN');
  const theme = useTheme();
  const navigate = useNavigate();

  const userData = useOutletContext();

  const [searchParams, setSearchParams] = useSearchParams();
  const auctionId = searchParams.get('id');
  const fullScreen = useMediaQuery(theme.breakpoints.down('md'));
  //  create a bid
  const [bidValue, setBidValue] = useState(auction.current_bid);
  const [isReady, setIsReady] = useState(false);

  // Users table
  const [page, setPage] = useState(0);

  const [order, setOrder] = useState('asc');

  const [selected, setSelected] = useState([]);

  const [orderBy, setOrderBy] = useState('name');

  const [rowsPerPage, setRowsPerPage] = useState(4);

  const handleChangePage = (event, newPage) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (event) => {
    setRowsPerPage(parseInt(event.target.value, 10));
    setPage(0);
  };

  const [userList, setUserList] = useState([]);
  const renewUserList = async () => {
    await fetch(`http://localhost:1009/auctee/all-bids/auction?id=${auction.Id}`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
    }).then((res) => {
      if (res.status === 200) {
        res.json().then((data) => {
          setUserList(data);
          setIsReady(true);
        });
      }
    });
    // eslint-disable-next-line no-unused-expressions
    bidValue <= auction.current_bid ? setBidValue(auction.current_bid) : setBidValue(bidValue);
  };
  const emptyRows = page > 0 ? Math.max(0, (1 + page) * rowsPerPage - userList?.length) : 0;

  const [socketUrl, setSocketUrl] = useState('ws://localhost:1009/auctee/ws');
  const [messageHistory, setMessageHistory] = useState([]);

  const { lastMessage, sendJsonMessage } = useWebSocket(socketUrl);
  const handleReady = async () => {
    if (currentInCome < auction.current_bid) {
      alert('Số dư trong ví không đủ để tham gia đấu giá');
      setIsReady(false);
    } else {
      setIsReady(true);
    }
  };

  const handleClickSendMessage = useCallback((body) => sendJsonMessage(body), []);
  const [errorMessage, setErrorMessage] = useState('');

  const handleBid = async () => {
    if (bidValue < auction.current_bid) {
      alert('Invalid value');
      return;
    }

    const payload = {
      bid_value: Number(bidValue),
      nickname: nickName,
    };
    await fetch(`http://localhost:1009/auctee/auction?auctionId=${auction.Id}&productId=${product.id}`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
      body: JSON.stringify(payload),
    }).then((res) => {
      if (res.status === 200) {
        setErrorMessage('');
        handleClickSendMessage({
          bid_value: bidValue,
          nickname: nickName,
          user_id: idPlayer,
          auction_id: parseInt(auctionId, 10),
        });
        setIsReady(true);
      } else {
        res.json().then((res) => {
          setBidValue(auction.current_bid);
          setErrorMessage(res.message);
        });
      }
    });
    // renewUserList();
  };

  const [nickName, setNickName] = useState('');
  const [idPlayer, setIdPlayer] = useState('');

  const [currentInCome, setCurrentInCome] = useState();

  const [isOrderCreated, setIsOrderCreated] = useState(false);
  const [isCheckedOut, setIsCheckedOut] = useState(false);
  const [paymentId, setPaymentId] = useState();
  const createPayment = async () => {
    if (
      (userList[0]?.user_id === idPlayer && isEnded) ||
      (userList[0]?.user_id === idPlayer &&
        (userList[0]?.bid_value >= product.expect_price || auction.current_bid >= product.expect_price))
    ) {
      await fetch(`http://localhost:1003/auctee/user/checkout/auction?id=${auction.Id}`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
      }).then((res) => {
        if (res.status === 200) {
          setIsOrderCreated(true);
          res.json().then((res) => {
            setPaymentId(res.id);
          });
        } else {
          res.json().then((res) => {
            if (res.message === 'order is pending') {
              setIsOrderCreated(true);
              setIsCheckedOut(true);
              setPaymentId(res.id);
            }
          });
        }
      });
    }
  };

  useEffect(() => {
    if (userData) {
      setNickName(userData.nickname);
      setCurrentInCome(userData.total_income);
      setIdPlayer(userData.id);
      if (userData.present_auction.toString() === auctionId) {
        setIsReady(true);
      }
    }
  }, [userData]);

  useEffect(() => {
    // eslint-disable-next-line no-unused-expressions
    const body = lastMessage && JSON.parse(lastMessage.data);
    // eslint-disable-next-line no-unused-expressions
    auctionId.length > 0 && body?.auction_id === Number(auctionId) && renewUserList();
  }, [lastMessage]);

  useEffect(() => {
    // eslint-disable-next-line no-unused-expressions
    lastMessage && setMessageHistory((prev) => prev.concat(lastMessage));
  }, [lastMessage, setMessageHistory]);

  const [isEnded, setIsEnd] = useState(false);

  // initial loading
  useEffect(() => {
    // eslint-disable-next-line no-unused-expressions
    (new Date(auction.end_time) < new Date() || auction.current_bid >= product.expect_price) && setIsEnd(true);
    // eslint-disable-next-line no-unused-expressions
    !isReady && renewUserList();
  }, [product, auction, userList]);

  useEffect(() => {
    // eslint-disable-next-line no-unused-expressions
    userList.length > 0 && createPayment();
  }, [userList]);

  return (
    <>
      <Dialog
        sx={{ margin: 'auto', maxWidth: '530px !important' }}
        BackdropProps={{
          style: { backgroundColor: 'rgba(0,0,30,0.2)' },
          invisible: true,
        }}
        fullScreen={fullScreen}
        open={isOrderCreated}
      >
        <Typography fontStyle="italic" variant="subtitle1" sx={{ px: 3, pt: 2, pb: 0 }}>
          {!isCheckedOut
            ? ' Chúc mừng bạn đã là người chốt thành công sản phẩm! Vui lòng hoàn tất thanh toán trễ nhất sau 3 ngày kể từ ngày đấu giá thành công'
            : 'Đơn hàng của bạn đã được tạo! Bạn có muốn kiểm tra thông tin đơn hàng ngay bây giờ?'}
        </Typography>
        <Box component="img" src="/static/congrate.svg" sx={{ px: '100px' }} />
        <Stack sx={{ p: 2, pt: 0 }} justifyContent="flex-end" direction="row" alignItems="center">
          <Button
            sx={{
              color: 'inherit',
              bgcolor: 'transparent',
              opacity: 0.85,
              border: '1px solid white',
              textTransform: 'none',
              '&:hover': {
                bgcolor: 'transparent',
                opacity: 1,
                border: '1px solid black',
              },
            }}
            onClick={() => setIsOrderCreated(false)}
          >
            Trở lại
          </Button>
          {isCheckedOut ? (
            // Go to detail page
            <Button
              disableRipple
              color="error"
              variant="contained"
              sx={{
                ml: 1,
                textTransform: 'none',
                color: 'white',
                bgcolor: '#f44336',
              }}
              onClick={() => {
                navigate(`/auctee/user/order/?id=${paymentId}`);
              }}
              autoFocus
            >
              Xem chi tiết đơn hàng
            </Button>
          ) : (
            // Go to page create order to add address
            <Button
              disableRipple
              color="error"
              variant="contained"
              sx={{
                ml: 1,
                textTransform: 'none',
                color: 'white',
                bgcolor: '#f44336',
              }}
              onClick={() => {
                navigate(`/auctee/user/order/?id=${paymentId}`);
              }}
              autoFocus
            >
              Thanh toán ngay
            </Button>
          )}
        </Stack>
      </Dialog>

      {/* Bidding area */}
      <Stack justifyContent="space-between" alignItems="center" direction="row" sx={{ mt: '30px !important' }}>
        <Stack direction="row" alignItems="center">
          <Typography variant="body1" noWrap>
            Bước giá: &nbsp;&nbsp;&nbsp;
          </Typography>
          <Typography sx={{ fontSize: '1.3rem', textOverflow: 'ellipsis' }} variant="body1" noWrap>
            {auction.price_per_step.toLocaleString('tr-TR', {
              style: 'currency',
              currency: 'VND',
            })}
          </Typography>
        </Stack>
        <Stack sx={{ color: `${theme.palette.background.main}` }} direction="row" alignItems="center">
          <EditButton
            disabled={isEnded}
            onClick={() => {
              if (bidValue > auction.current_bid) {
                setBidValue(bidValue - auction.price_per_step);
              }
            }}
          >
            -
          </EditButton>
          <input
            value={bidValue || ''}
            onChange={(e) => setBidValue(e.target.value)}
            cursor="text"
            disabled="true"
            style={{
              textAlign: 'center',
              lineHeight: 0,
              height: '40px',
              width: '80px',
              fontSize: '1.2rem',
              color: 'inherit',
              border: `1px solid ${theme.palette.background.main}`,
            }}
            step={`${auction.price_per_step}`}
            type="number"
            max={`${product.expect_price}`}
            min={`${auction.current_bid}`}
          />
          <EditButton
            disabled={isEnded}
            onClick={() => {
              if (bidValue < product.expect_price) {
                setBidValue(bidValue + auction.price_per_step);
              }
            }}
          >
            +
          </EditButton>
        </Stack>
        <Stack>
          {isReady ? (
            <StyledButton disabled={isEnded} onClick={() => handleBid()}>
              Đấu Giá <Icon style={{ fontSize: '1.2rem', marginLeft: 2 }} icon="mingcute:auction-line" />
            </StyledButton>
          ) : (
            <StyledButton
              sx={{
                bgcolor: 'transparent',
                color: theme.palette.background.main,
                border: `1px solid ${theme.palette.background.main}`,
                opacity: 1,
              }}
              onClick={() => handleReady()}
            >
              Tham gia
            </StyledButton>
          )}
        </Stack>
      </Stack>
      <Stack sx={{ height: '24px' }}>
        {errorMessage.length > 0 && (
          <Typography sx={{ position: 'relative', ml: 2 }} variant="subtitle2" color="error">
            <Icon style={{ position: 'absolute', left: '-15px', bottom: '5px' }} icon="bi:exclamation-circle" /> &nbsp;
            {errorMessage}
          </Typography>
        )}
      </Stack>
      {/* All users */}
      <Container sx={{ px: '0 !important', bgcolor: 'transparent' }}>
        <Card
          sx={{
            bgcolor: alpha(theme.palette.background.main, 0.05),
            borderRadius: 1,
            border: `1px solid ${theme.palette.background.main}`,
          }}
        >
          <Scrollbar sx={{ maxHeight: '300px' }}>
            <TableContainer sx={{ minWidth: '100%' }}>
              <Table>
                <UserListHead
                  order={order}
                  orderBy={orderBy}
                  headLabel={TABLE_HEAD}
                  rowCount={userList.length}
                  numSelected={selected.length}
                />
                <TableBody>
                  {userList.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage).map((row, index) => {
                    const user = {
                      userId: row.user_id,
                      nickName: row.nickname || 'unknown player',
                      bidValue: row.bid_value,
                      bidTime: timeAgo.format(new Date(row.bid_time)),
                    };
                    const isItemSelected = selected.indexOf(nickName) !== -1;

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
                                color: `${user.userId === idPlayer ? theme.palette.background.main : 'inherit'}`,
                              }}
                              variant="subtitle2"
                              noWrap
                            >
                              {user.nickName}
                              {index === 0 && (
                                <Icon
                                  style={{
                                    color: '#ebab09',
                                    marginLeft: 1.5,
                                    marginBottom: '-3px',
                                    fontSize: '1.2rem',
                                  }}
                                  icon="ph:crown-simple-bold"
                                />
                              )}
                            </Typography>
                            {user.userId === idPlayer && (
                              <Typography variant="body2" sx={{ fontStyle: 'italic' }}>
                                (Bạn)
                              </Typography>
                            )}
                          </Stack>
                        </TableCell>
                        <TableCell align="left">
                          {user.bidValue.toLocaleString('tr-TR', {
                            style: 'currency',
                            currency: 'VND',
                          })}
                        </TableCell>
                        <TableCell align="left">{user.bidTime}</TableCell>
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
            count={userList.length}
            rowsPerPage={rowsPerPage}
            page={page}
            onPageChange={handleChangePage}
            onRowsPerPageChange={handleChangeRowsPerPage}
          />
        </Card>
      </Container>
    </>
  );
};

export default BidSection;
