import { useState, useEffect, lazy, useContext, Suspense } from 'react';
import { Link as RouterLink, useNavigate, useLocation, useSearchParams } from 'react-router-dom';
import { useForm } from 'react-hook-form';
import Select from 'react-select';
import moment from 'moment';

// material
import {
  Typography,
  Stack,
  Button,
  Divider,
  Dialog,
  Checkbox,
  DialogTitle,
  RadioGroup,
  Radio,
  alertTitleClasses,
  FormControlLabel,
} from '@mui/material';

import { LoadingButton } from '@mui/lab';
import { Icon } from '@iconify/react';
import { styled, useTheme } from '@mui/material/styles';
import { Box } from '@mui/system';

import { ReloadContext } from '../../utils/Context';
import CountDown from '../../utils/countdown';
import UpdateAuctionForm from './UpdateAuctionForm';

export default function AuctionList() {
  const theme = useTheme();
  const navigate = useNavigate();
  const location = useLocation();

  const { isReloading, setIsReloading } = useContext(ReloadContext);

  const [isFetching, setIsFetching] = useState(true);

  // User information
  const [auctionsData, setAuctionData] = useState();

  // Get user's data base on access_token
  const handleFetchAuctionData = async () => {
    await fetch('http://localhost:1009/auctee/user/auctions', {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
    }).then((res) => {
      if (res.status === 200) {
        res.json().then((data) => {
          setAuctionData(data);
          setIsFetching(false);
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
  const handleDelete = async (id) => {
    await fetch(`http://localhost:1009/auctee/user/auction/detail?id=${id}`, {
      method: 'DELETE',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
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

  // Update default address ONLY
  const handleUpdateDefault = async (address) => {
    const payload = {
      firstname: address.firstname,
      lastname: address.lastname,
      phone: address.phone,
      province: address.province,
      district: address.district,
      email: address.email,
      sub_district: address.sub_district,
      address: address.address,
      type_address: address.type_address,
      is_default: true,
    };
    await fetch(`http://localhost:1001/auctee/user/address?id=${address.ID}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
      credentials: 'include',
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

  useEffect(() => {
    setIsReloading(false);
    handleFetchAuctionData();
  }, [isFetching, isReloading]);
  return !isFetching
    ? auctionsData?.map((auction, index) => (
        <Stack key={index}>
          {index === 0 && (
            <Typography sx={{ mb: 2, fontWeight: 500 }}>
              Hiện có: &nbsp;&nbsp;&nbsp; {auctionsData.length} phiên đấu giá
            </Typography>
          )}
          <Stack justifyContent="space-between" alignItems="flex-start" direction="row" sx={{ pb: 3 }}>
            {/* Left infor */}
            <Stack sx={{ width: '100%' }}>
              <Stack alignItems="center" direction="row" sx={{ width: '70%' }}>
                <Typography variant="subtitle1" sx={{ color: 'inherit' }}>
                  ID sản phẩm : {auction.product_id}
                </Typography>
                <Stack sx={{ ml: 2, pl: 2, borderLeft: '2px solid grey' }}>
                  <Typography variant="body1" sx={{ color: 'inherit' }}>
                    {auction.name}
                  </Typography>
                </Stack>
              </Stack>
              <Stack direction="row">
                <Stack sx={{ flex: 0.3 }}>
                  <Stack direction="row">
                    <Typography fontSize={'0.9rem'} variant="body2">
                      Trạng thái &nbsp;&nbsp;:
                    </Typography>
                    <Typography
                      fontSize={'0.85rem'}
                      variant="body2"
                      sx={{ color: `${auction.is_active ? 'green' : 'red'}` }}
                    >
                      &nbsp;
                      {auction.is_active ? 'Đang diễn ra' : 'Chưa kích hoạt'}
                    </Typography>
                  </Stack>
                  <Typography fontSize={'0.9rem'} variant="body2" sx={{ color: 'inherit' }}>
                    Bước giá &nbsp;&nbsp;&nbsp; :&nbsp;{' '}
                    {auction.price_per_step.toLocaleString('tr-TR', {
                      style: 'currency',
                      currency: 'VND',
                    })}
                  </Typography>
                  <Typography fontSize={'0.9rem'} variant="body2" sx={{ color: 'green' }}>
                    Giá hiện tại :&nbsp;
                    {auction.current_bid.toLocaleString('tr-TR', {
                      style: 'currency',
                      currency: 'VND',
                    })}
                  </Typography>
                </Stack>
                <Stack justifyContent="flex-end" sx={{ flex: 0.7 }}>
                  <Stack
                    fontSize={'0.9rem'}
                    variant="body2"
                    direction="row"
                    sx={{
                      color: 'inherit',
                      maxHeight: '22.63px',
                    }}
                  >
                    Ngày bắt đầu &nbsp; :&nbsp;&nbsp;{new Date(moment(auction.start_time)).toLocaleTimeString('en-US')}{' '}
                    {new Date(moment(auction.start_time)).toLocaleDateString('en-GB')}
                  </Stack>
                  <Stack
                    fontSize={'0.9rem'}
                    variant="body2"
                    direction="row"
                    sx={{
                      color: 'inherit',
                      maxHeight: '22.63px',
                    }}
                  >
                    Kết thúc trong :&nbsp;
                    <CountDown time={auction.end_time} />
                  </Stack>
                </Stack>
              </Stack>
            </Stack>
            {/* Right button */}
            {!auction.is_active && (
              <Stack alignItems="flex-end" sx={{ width: '20%' }}>
                {/* Update and Delete Address */}
                <UpdateAuctionForm
                  key={auction.Id}
                  auction={auction}
                  handleDelete={handleDelete}
                  handleUpdateDefault={handleUpdateDefault}
                />
              </Stack>
            )}
          </Stack>
          {auctionsData.length - index !== 1 && <Divider sx={{ mb: 2 }} />}
        </Stack>
      ))
    : auctionsData?.length && <>Loading...</>;
}
