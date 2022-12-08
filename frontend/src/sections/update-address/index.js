import { useState, useEffect, lazy, useContext, Suspense } from 'react';
import { Link as RouterLink, useNavigate, useLocation, useSearchParams } from 'react-router-dom';
import { useForm } from 'react-hook-form';
import Select from 'react-select';
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

import UpdateAddressForm from './UpdateAddressForm';

export default function AddressList() {
  const theme = useTheme();
  const navigate = useNavigate();
  const location = useLocation();

  const { isReloading, setIsReloading } = useContext(ReloadContext);

  const [isFetching, setIsFetching] = useState(false);

  // User information
  const [addressesData, setAddressesData] = useState();

  // Get user's data base on access_token
  const handleFetchAddressData = async () => {
    await fetch('http://localhost:8080/auctee/user/addresses', {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',

      mode: 'cors',
    }).then((res) => {
      if (res.status === 200) {
        res.json().then((data) => {
          setAddressesData(data);
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

  // Delete address
  const handleDelete = async (id) => {
    await fetch(`http://localhost:8080/auctee/user/address?id=${id}`, {
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
    await fetch(`http://localhost:8080/auctee/user/address?id=${address.ID}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
      credentials: 'include',
      mode: 'cors',
    }).then((res) => {
      if (res.status === 200) {
        setIsReloading(true);
        setIsFetching(false);
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
    // eslint-disable-next-line no-unused-expressions
    !isFetching && handleFetchAddressData();
  }, [isFetching, isReloading]);
  return addressesData?.length > 0 ? (
    addressesData.map((address, index) => (
      <Stack key={index}>
        <Stack justifyContent="space-between" alignItems="center" direction="row" sx={{ pb: 3 }}>
          {/* Left infor */}
          <Stack sx={{ width: '100%' }}>
            <Stack alignItems="center" direction="row" sx={{ width: '40%' }}>
              <Typography fontSize={'1rem'} variant="body2" sx={{ color: 'black' }}>
                {address.lastname} {address.firstname}
              </Typography>
              <Stack sx={{ ml: 2, pl: 2, borderLeft: '1px solid grey' }}>
                <Typography fontSize={'0.9rem'} variant="caption" sx={{ color: 'inherit' }}>
                  (+84) &nbsp;{address.phone.slice(1)}
                </Typography>
              </Stack>
            </Stack>
            <Typography fontSize={'0.9rem'} variant="body2" sx={{ color: 'black', opacity: 0.6 }}>
              {address.address}
            </Typography>
            <Typography fontSize={'0.9rem'} variant="body2" sx={{ color: 'black', opacity: 0.6 }}>
              {address.sub_district}, {address.district}, {address.province}
            </Typography>
            {address.is_default && (
              <Button
                fontSize="0.1rem"
                sx={{
                  borderRadius: 0,
                  color: theme.palette.background.main,
                  textTransform: 'none',
                  width: '10%',
                  px: 0.5,
                  mt: 0.5,
                  py: 0,
                  fontWeight: 500,
                  border: `1px solid ${theme.palette.background.main}`,
                }}
              >
                Mặc định
              </Button>
            )}
          </Stack>
          {/* Right button */}
          <Stack alignItems="flex-end" sx={{ width: '20%' }}>
            {/* Update and Delete Address */}
            <UpdateAddressForm
              key={address.ID}
              address={address}
              handleDelete={handleDelete}
              handleUpdateDefault={handleUpdateDefault}
            />
          </Stack>
        </Stack>
        {addressesData.length - index !== 1 && <Divider sx={{ mb: 2 }} />}
      </Stack>
    ))
  ) : (
    <>Có lỗi xảy ra</>
  );
}
