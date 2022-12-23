import { useState, useEffect, lazy, useContext, Suspense } from 'react';
import { Link, useNavigate, useLocation, useSearchParams } from 'react-router-dom';
import { useForm } from 'react-hook-form';
import Select from 'react-select';
import moment from 'moment';

// material
import {
  Typography,
  Stack,
  Button,
  Divider,
  Avatar,
  ClickAwayListener,
} from '@mui/material';

import { LoadingButton } from '@mui/lab';
import { Icon } from '@iconify/react';
import { styled, useTheme } from '@mui/material/styles';
import { Box } from '@mui/system';
import Iconify from '../../components/Iconify';
import { ReloadContext } from '../../utils/Context';
import CountDown from '../../utils/countdown';
import UpdateAuctionForm from './UpdateAuctionForm';

export default function AuctionList() {
  const theme = useTheme();
  const navigate = useNavigate();
  const location = useLocation();

  const { isReloading, setIsReloading } = useContext(ReloadContext);

  const [isFetching, setIsFetching] = useState(true);
  const [openFilter, setOpenFilter] = useState(false);
  const handleClick = () => {
    setOpenFilter((prev) => !prev);
  };

  const handleClickAway = () => {
    setOpenFilter(false);
  };
  // User information
  const [auctionsData, setAuctionData] = useState();

  const handleFetchAuctionData = async () => {
    await fetch('http://localhost:8080/auctee/user/auctions', {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',

      mode: 'cors',
    }).then((res) => {
      if (res.status === 200) {
        res.json().then((data) => {
          setAuctionData(data);
          setIsFetching(false);
          setFilteredAuctions(data);
        });
      }
      if (res.status === 401) {
        setIsFetching(true);
        alert('You need to login first');
        navigate('/auctee/login', { replace: true });
      }
    });
  };

  // Delete
  const handleDelete = async (auction) => {
    await fetch(`http://localhost:8080/auctee/user/auction/detail?id=${auction.Id}`, {
      method: 'DELETE',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',

      mode: 'cors',
    }).then((res) => {
      if (res.status === 200) {
        setIsReloading(true);
      }
      if (res.status === 401) {
        setIsReloading(false);
        alert('You need to login first');
        setIsFetching(true);
        navigate('/auctee/login', { replace: true });
      }
    });
  };

  // filtered auction
  const [filteredAuctions, setFilteredAuctions] = useState([]);

  // 1" chưa kích hoạt 2: đang diễn ra 3: kết thúc 4: chờ thanh toán
  const handleFilter = (type) => {
    setFilteredAuctions([]);
    if (auctionsData) {
      switch (type) {
        case 1:
          auctionsData.forEach((auction) => {
            if (!auction.is_active && new Date(auction.end_time) > new Date()) {
              setFilteredAuctions((current) => [...current, auction]);
            }
          });
          break;
        case 2:
          auctionsData.forEach((auction) => {
            if (new Date(auction.end_time) > new Date() && auction.is_active) {
              setFilteredAuctions((current) => [...current, auction]);
            }
          });
          break;
        case 3:
          auctionsData.forEach((auction) => {
            if (new Date(auction.end_time) < new Date()) {
              setFilteredAuctions((current) => [...current, auction]);
            }
          });
          break;
        case 4:
          auctionsData.forEach((auction) => {
            if (new Date(auction.end_time) < new Date() && auction.winner_id > 0) {
              setFilteredAuctions((current) => [...current, auction]);
            }
          });
          break;
        default:
          setFilteredAuctions(auctionsData);
          break;
      }
    }
    setOpenFilter(false);
  };

  useEffect(() => {
    // eslint-disable-next-line no-unused-expressions
    !auctionsData && handleFetchAuctionData();
  }, [auctionsData]);

  useEffect(() => {
    setIsReloading(false);
    // eslint-disable-next-line no-unused-expressions
    isReloading && handleFetchAuctionData();
  }, [isFetching, isReloading]);

  return (
    <>
      <Stack sx={{ pb: 1 }} direction="row" justifyContent="space-between" alignItems="center">
        <Typography sx={{ fontWeight: 500 }}>
          Hiện có: &nbsp;&nbsp;&nbsp; {auctionsData?.length} phiên đấu giá
        </Typography>
        <ClickAwayListener onClickAway={handleClickAway}>
          <Stack sx={{ position: 'relative' }}>
            <Button
              sx={{ fontSize: '1rem', fontWeight: 500 }}
              disableRipple
              color="inherit"
              endIcon={<Iconify style={{ fontSize: '1.4rem' }} icon="ic:round-filter-list" />}
              onClick={handleClick}
            >
              Lọc&nbsp;
            </Button>
            {openFilter && (
              <Stack
                sx={{
                  boxShadow: 2,
                  borderRadius: 0.5,
                  fontSize: '0.9rem',
                  py: 0.5,
                  zIndex: 999,
                  bgcolor: 'white',
                  position: 'absolute',
                  bottom: '-135px',
                  right: 0,
                  minWidth: '100%',
                }}
              >
                {/* All  */}
                <Button
                  onClick={() => handleFilter(0)}
                  sx={{
                    fontWeight: 400,
                    justifyContent: 'flex-end',
                    whiteSpace: 'nowrap',
                    textTransform: 'none',
                    width: '100%',
                    py: 0.1,
                    pl: 2,
                    '&:hover': { opacity: 0.75 },
                  }}
                >
                  Tất cả
                </Button>
                {/* Inactive */}
                <Button
                  onClick={() => handleFilter(1)}
                  sx={{
                    fontWeight: 400,
                    justifyContent: 'flex-end',
                    whiteSpace: 'nowrap',
                    textTransform: 'none',
                    width: '100%',
                    py: 0.1,
                    pl: 2,
                    color: 'inherit',
                    opacity: 0.7,
                    '&:hover': { opacity: 0.55 },
                  }}
                >
                  Chưa kích hoạt
                </Button>
                {/* Started */}
                <Button
                  onClick={() => handleFilter(2)}
                  sx={{
                    fontWeight: 400,
                    justifyContent: 'flex-end',
                    whiteSpace: 'nowrap',
                    textTransform: 'none',
                    width: '100%',
                    py: 0.1,
                    pl: 2,
                    color: 'green',
                    '&:hover': { opacity: 0.75 },
                  }}
                >
                  Đang diễn ra
                </Button>
                {/* Ended */}
                <Button
                  onClick={() => handleFilter(3)}
                  sx={{
                    fontWeight: 400,
                    justifyContent: 'flex-end',
                    whiteSpace: 'nowrap',
                    textTransform: 'none',
                    width: '100%',
                    pl: 4,
                    py: 0.1,
                    color: 'red',
                    '&:hover': { opacity: 0.75 },
                  }}
                >
                  Kết thúc
                </Button>
                {/* Pending payment */}
                <Button
                  onClick={() => handleFilter(4)}
                  sx={{
                    fontWeight: 400,
                    justifyContent: 'flex-end',
                    whiteSpace: 'nowrap',
                    textTransform: 'none',
                    width: '100%',
                    pl: 2,
                    py: 0.1,
                    color: '#f5b70c',
                    '&:hover': { opacity: 0.75 },
                  }}
                >
                  Chờ thanh toán
                </Button>
              </Stack>
            )}
          </Stack>
        </ClickAwayListener>
      </Stack>
      {!isFetching ? (
        filteredAuctions?.map((auction, index) => (
          <Stack key={index}>
            <Stack justifyContent="space-between" alignItems="flex-start" direction="row" sx={{ pb: 3 }}>
              {/* Left infor */}
              <Stack sx={{ width: '100%' }}>
                <Stack alignItems="center" direction="row" sx={{ width: '70%' }}>
                  <Typography variant="subtitle1" sx={{ color: 'inherit' }}>
                    ID sản phẩm : {auction.product_id}
                  </Typography>
                  <Stack sx={{ ml: 2, pl: 2, borderLeft: '2px solid grey' }}>
                    <Link
                      to={`/auctee/auction/detail?id=${auction.Id}&product=${auction.product_id}`}
                      style={{
                        color: 'inherit',
                        textDecoration: 'none',
                      }}
                    >
                      {auction.name}
                    </Link>
                  </Stack>
                </Stack>

                <Stack direction="row">
                  <Stack justifyContent="flex-end" sx={{ flex: 0.3 }}>
                    <Stack direction="row">
                      <Typography fontSize={'0.9rem'} variant="body2">
                        Trạng thái &nbsp;&nbsp;:
                      </Typography>
                      {new Date(auction.end_time) < new Date() ? (
                        <Typography
                          fontWeight={600}
                          fontSize={'0.85rem'}
                          variant="body2"
                          sx={{ color: 'red', opacity: 0.8 }}
                        >
                          &nbsp; Kết thúc
                        </Typography>
                      ) : (
                        <Typography
                          fontWeight={500}
                          fontSize={'0.85rem'}
                          variant="body2"
                          sx={{ color: `${auction.is_active ? 'green' : 'grey'}`, opacity: 0.8 }}
                        >
                          &nbsp;
                          {auction.is_active ? 'Đang diễn ra' : 'Chưa kích hoạt'}
                        </Typography>
                      )}
                    </Stack>
                    <Typography fontSize={'0.9rem'} variant="body2" sx={{ color: 'inherit' }}>
                      Bước giá &nbsp;&nbsp;&nbsp; :&nbsp;{' '}
                      {auction.price_per_step.toLocaleString('tr-TR', {
                        style: 'currency',
                        currency: 'VND',
                      })}
                    </Typography>
                    <Typography fontSize={'0.9rem'} variant="body2">
                      Giá hiện tại :&nbsp;
                      {auction.current_bid.toLocaleString('tr-TR', {
                        style: 'currency',
                        currency: 'VND',
                      })}
                    </Typography>
                  </Stack>
                  <Stack justifyContent="flex-end" sx={{ flex: 0.5, mb: 0.2 }}>
                    <Stack
                      fontSize={'0.9rem'}
                      variant="body2"
                      direction="row"
                      sx={{
                        color: 'inherit',
                        maxHeight: '22.63px',
                      }}
                    >
                      {new Date(auction.end_time) < new Date() ? (
                        <Stack alignItems="flex-end" sx={{ fontStyle: 'italic' }} direction="row">
                          Chú thích: &nbsp;&nbsp;
                          <Stack color={auction.winner_id ? '#f5b70c' : 'red'} variant="body2">
                            {auction.winner_id ? 'Chờ thanh toán' : 'Không có người tham gia'}&nbsp;&nbsp;&nbsp;&nbsp;
                          </Stack>
                          {auction.winner_id ? (
                            <Link
                              to={`/auctee/auction/detail?id=${auction.Id}&product=${auction.product_id}`}
                              variant="body1"
                              style={{
                                fontSize: '0.8rem',
                                color: '#2c8df5',
                              }}
                            >
                              Chi tiết
                            </Link>
                          ) : (
                            ''
                          )}
                        </Stack>
                      ) : (
                        ''
                      )}
                    </Stack>
                    <Typography
                      fontSize={'0.9rem'}
                      variant="body2"
                      direction="row"
                      sx={{
                        color: 'inherit',
                      }}
                    >
                      Ngày bắt đầu &nbsp; :&nbsp;&nbsp;
                      {new Date(moment(auction.start_time)).toLocaleTimeString('en-US')} &nbsp;
                      {new Date(moment(auction.start_time)).toLocaleDateString('en-GB')}
                    </Typography>
                    <Stack
                      fontSize={'0.9rem'}
                      variant="body2"
                      direction="row"
                      sx={{
                        color: 'inherit',
                      }}
                    >
                      Kết thúc trong :&nbsp;
                      <CountDown time={auction.end_time} />
                    </Stack>
                  </Stack>
                  <Avatar
                    height={100}
                    sx={{ flex: 0.1, height: 68 }}
                    variant="square"
                    alt={auction.image_path}
                    src={auction.image_path}
                  />
                </Stack>
              </Stack>
              {/* Right button */}
              {/* {!auction.is_active && ( */}
              <Stack alignItems="flex-end" sx={{ width: '20%' }}>
                {/* Update and Delete Address */}
                <UpdateAuctionForm index={index} auction={auction} handleDelete={handleDelete} />
              </Stack>
              {/* )} */}
            </Stack>
            {auctionsData.length - index !== 1 && <Divider sx={{ mb: 2 }} />}
          </Stack>
        ))
      ) : (
        <>Loading...</>
      )}
    </>
  );
}
