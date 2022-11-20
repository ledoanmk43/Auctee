import { useState, useEffect, useCallback } from 'react';
import { Link as RouterLink, useNavigate, useSearchParams, useLocation } from 'react-router-dom';
import useWebSocket, { ReadyState } from 'react-use-websocket';
import TimeAgo from 'javascript-time-ago';
// English.
import vi from 'javascript-time-ago/locale/vi';
import { Icon } from '@iconify/react';
import {
  Card,
  Table,
  Stack,
  Avatar,
  Button,
  Checkbox,
  TableRow,
  TableBody,
  TableCell,
  Container,
  Typography,
  TableContainer,
  TablePagination,
} from '@mui/material';
import { styled, useTheme, alpha } from '@mui/material/styles';
import moment from 'moment';
import 'moment/locale/vi';
import Label from '../../components/Label';
import Scrollbar from '../../components/Scrollbar';
import SearchNotFound from '../../components/SearchNotFound';
import { UserListHead, UserListToolbar, UserMoreMenu } from '../@dashboard/user';

const TABLE_HEAD = [
  { id: 'name', label: 'Người tham gia', alignRight: false },
  { id: 'company', label: 'Số tiền ', alignRight: false },
  { id: 'role', label: 'Thời gian ', alignRight: false },
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
  const [searchParams, setSearchParams] = useSearchParams();
  const auctionId = searchParams.get('id');
  //  create a bid
  const [bidValue, setBidValue] = useState();
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
        });
      }
    });
    setIsReloading(true);
  };
  const emptyRows = page > 0 ? Math.max(0, (1 + page) * rowsPerPage - userList?.length) : 0;

  const [socketUrl, setSocketUrl] = useState('ws://localhost:1009/auctee/ws');
  const [messageHistory, setMessageHistory] = useState([]);

  const { lastMessage, readyState, sendJsonMessage } = useWebSocket(socketUrl);
  const handleReady = async () => {
    if (currentInCome < auction.current_bid) {
      alert('Số dư trong ví không đủ để tham gia đấu giá');
      setIsReady(false);
    } else {
      handleClickSendMessage();
      setIsReady(true);
    }
  };

  const handleClickSendMessage = useCallback((body) => sendJsonMessage(body), []);

  const [errorMessage, setErrorMessage] = useState('');

  const [isReloading, setIsReloading] = useState(false);
  const handleBid = async () => {
    if (bidValue < auction.current_bid) {
      alert('Invalid value');
      return;
    }

    const payload = {
      bid_value: bidValue,
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
        });
        setIsReady(true);
      } else {
        res.json().then((res) => {
          setBidValue(auction.current_bid);
          setErrorMessage(res.message);
        });
      }
    });
    renewUserList();
  };

  const [nickName, setNickName] = useState('');
  const [idPlayer, setIdPlayer] = useState('');

  const [currentInCome, setCurrentInCome] = useState();
  const fetchUser = async () => {
    await fetch(`http://localhost:1001/auctee/user/profile`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
    }).then((res) => {
      if (res.status === 200) {
        res.json().then((data) => {
          setNickName(data.nickname);
          setCurrentInCome(data.total_income);
          setIdPlayer(data.id);
          if (data.present_auction.toString() === auctionId) {
            setIsReady(true);
          }
        });
      }
    });
  };

  useEffect(() => {
    // eslint-disable-next-line no-unused-expressions
    nickName.length === 0 && fetchUser();
  }, [nickName]);

  useEffect(() => {
    if (lastMessage !== null) {
      setMessageHistory((prev) => prev.concat(lastMessage));
      renewUserList();
    }
  }, [isReady, lastMessage, setMessageHistory, isReloading]);
  useEffect(() => {
    setBidValue(auction.current_bid);
  }, [product, auction]);

  return (
    <>
      {/* Bidding area */}
      <Stack justifyContent="space-between" alignItems="center" direction="row" sx={{ mt: '30px !important' }}>
        <Stack direction="row">
          <Typography sx={{ fontSize: '1.2rem', textOverflow: 'ellipsis' }} variant="body1" noWrap>
            Bước giá: &nbsp;
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
            disabled={!isReady}
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
            style={{
              textAlign: 'center',
              lineHeight: 0,
              height: '40px',
              width: '80px',
              fontSize: '1.2rem',
              color: 'inherit',
              border: `1px solid ${theme.palette.background.main}`,
            }}
            type="number"
            max={`${product.expect_price}`}
            min={`${auction.current_bid}`}
          />
          <EditButton
            disabled={!isReady}
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
            <StyledButton onClick={() => handleBid()}>
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
