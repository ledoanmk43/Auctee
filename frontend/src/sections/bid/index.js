import { useState, useEffect, lazy, Suspense } from 'react';
import { Link as RouterLink, useNavigate, useSearchParams, useLocation } from 'react-router-dom';

import { filter } from 'lodash';
import { sentenceCase } from 'change-case';
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
import { styled, useTheme } from '@mui/material/styles';
import Label from '../../components/Label';
import Scrollbar from '../../components/Scrollbar';
import SearchNotFound from '../../components/SearchNotFound';
import { UserListHead, UserListToolbar, UserMoreMenu } from '../@dashboard/user';
// mock
import USERLIST from '../../API/user';

const TABLE_HEAD = [
  { id: 'name', label: 'Người tham gia', alignRight: false },
  { id: 'company', label: 'Số tiền ', alignRight: false },
  { id: 'role', label: 'Thời gian ', alignRight: false },
  { id: 'isVerified', label: 'Verified', alignRight: false },
  { id: 'status', label: 'Status', alignRight: false },
];

function descendingComparator(a, b, orderBy) {
  if (b[orderBy] < a[orderBy]) {
    return -1;
  }
  if (b[orderBy] > a[orderBy]) {
    return 1;
  }
  return 0;
}

function getComparator(order, orderBy) {
  return order === 'desc'
    ? (a, b) => descendingComparator(a, b, orderBy)
    : (a, b) => -descendingComparator(a, b, orderBy);
}

function applySortFilter(array, comparator, query) {
  const stabilizedThis = array.map((el, index) => [el, index]);
  stabilizedThis.sort((a, b) => {
    const order = comparator(a[0], b[0]);
    if (order !== 0) return order;
    return a[1] - b[1];
  });
  if (query) {
    return filter(array, (_user) => _user.name.toLowerCase().indexOf(query.toLowerCase()) !== -1);
  }
  return stabilizedThis.map((el) => el[0]);
}

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
const BidSection = ({ product, auction }) => {
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

  const [filterName, setFilterName] = useState('');

  const [rowsPerPage, setRowsPerPage] = useState(5);

  const handleRequestSort = (event, property) => {
    const isAsc = orderBy === property && order === 'asc';
    setOrder(isAsc ? 'desc' : 'asc');
    setOrderBy(property);
  };

  const handleChangePage = (event, newPage) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (event) => {
    setRowsPerPage(parseInt(event.target.value, 10));
    setPage(0);
  };

  const emptyRows = page > 0 ? Math.max(0, (1 + page) * rowsPerPage - USERLIST.length) : 0;

  const filteredUsers = applySortFilter(USERLIST, getComparator(order, orderBy), filterName);

  const isUserNotFound = filteredUsers.length === 0;

  const handleReady = async () => {
    if (presentAuction.toString() !== auctionId && presentAuction > 0) {
      alert('Bạn đang tham gia một cuộc đấu giá khác');
      setIsReady(false);
      return;
    }
    const payload = {
      present_auction: auctionId,
    };
    await fetch(`http://localhost:1001/auctee/user/profile/setting`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
      credentials: 'include',
    }).then((res) => {
      if (res.status === 200) {
        setIsReady(true);
      } else {
        setIsReady(false);
        alert('Something goes wrong. Please try again');
      }
    });
  };

  const [presentAuction, setPresentAuction] = useState();
  const fetchUser = async () => {
    await fetch(`http://localhost:1001/auctee/user/profile`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
    }).then((res) => {
      if (res.status === 200) {
        res.json().then((data) => {
          setPresentAuction(Number(data.present_auction));
          if (data.present_auction.toString() === auctionId) {
            setIsReady(true);
          }
        });
      }
    });
  };

  useEffect(() => {
    fetchUser();
    setBidValue(auction.current_bid);
  }, [product, auction, isReady]);
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
            <StyledButton onClick={() => {}}>
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
              Sẵn sàng
            </StyledButton>
          )}
        </Stack>
      </Stack>
      {/* All users */}
      <Container sx={{ px: '0 !important' }}>
        <Card>
          <Scrollbar sx={{ maxHeight: '300px' }}>
            <TableContainer sx={{ minWidth: '100%' }}>
              <Table>
                <UserListHead
                  order={order}
                  orderBy={orderBy}
                  headLabel={TABLE_HEAD}
                  rowCount={USERLIST.length}
                  numSelected={selected.length}
                  onRequestSort={handleRequestSort}
                />
                <TableBody>
                  {filteredUsers.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage).map((row) => {
                    const { id, name, role, status, company, avatarUrl, isVerified } = row;
                    const isItemSelected = selected.indexOf(name) !== -1;

                    return (
                      <TableRow
                        hover
                        key={id}
                        tabIndex={-1}
                        role="checkbox"
                        selected={isItemSelected}
                        aria-checked={isItemSelected}
                      >
                        <TableCell component="th" scope="row" padding="none" sx={{ pl: 2 }}>
                          <Stack direction="row" alignItems="center" spacing={2}>
                            <Avatar alt={name} src={avatarUrl} />
                            <Typography variant="subtitle2" noWrap>
                              {name}
                            </Typography>
                          </Stack>
                        </TableCell>
                        <TableCell align="left">{company}</TableCell>
                        <TableCell align="left">{role}</TableCell>
                        <TableCell align="left">{isVerified ? 'Yes' : 'No'}</TableCell>
                        <TableCell align="left">
                          <Label variant="ghost" color={(status === 'banned' && 'error') || 'success'}>
                            {sentenceCase(status)}
                          </Label>
                        </TableCell>

                        <TableCell align="right">
                          <UserMoreMenu />
                        </TableCell>
                      </TableRow>
                    );
                  })}
                  {emptyRows > 0 && (
                    <TableRow style={{ height: 53 * emptyRows }}>
                      <TableCell colSpan={6} />
                    </TableRow>
                  )}
                </TableBody>

                {isUserNotFound && (
                  <TableBody>
                    <TableRow>
                      <TableCell align="center" colSpan={6} sx={{ py: 3 }}>
                        <SearchNotFound searchQuery={filterName} />
                      </TableCell>
                    </TableRow>
                  </TableBody>
                )}
              </Table>
            </TableContainer>
          </Scrollbar>

          <TablePagination
            rowsPerPageOptions={[5, 10, 25]}
            component="div"
            count={USERLIST.length}
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
