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

import UpdateProductForm from './UpdateProductForm';

export default function ProductList() {
  const theme = useTheme();
  const navigate = useNavigate();
  const location = useLocation();

  const { isReloading, setIsReloading } = useContext(ReloadContext);

  const [isFetching, setIsFetching] = useState(true);

  const [productsData, setProductData] = useState();

  // Get user's data base on access_token
  const handleFetchProductData = async () => {
    await fetch('http://localhost:1002/auctee/products', {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
    }).then((res) => {
      if (res.status === 200) {
        res.json().then((data) => {
          setProductData(data);
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
    await fetch(`http://localhost:1002/auctee/user/product/detail?id=${id}`, {
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
  }, [isFetching, isReloading]);

  useEffect(() => {
    // eslint-disable-next-line no-unused-expressions
    if (!productsData || isReloading) {
      handleFetchProductData();
      setIsReloading(false);
    }
  }, [productsData, isReloading]);

  return !isFetching ? (
    productsData?.map((product, index) => (
      <Stack key={index}>
        {index === 0 && (
          <Typography sx={{ mb: 2, fontWeight: 500 }}>
            Tổng kho: &nbsp;&nbsp;&nbsp; {productsData.length} sản phẩm
          </Typography>
        )}
        <Stack justifyContent="space-between" alignItems="flex-start" direction="row" sx={{ pb: 3 }}>
          {/* Left infor */}
          <Stack sx={{ width: '100%' }}>
            <Stack alignItems="center" direction="row" sx={{ width: '70%' }}>
              <Typography variant="subtitle1" sx={{ color: 'inherit' }}>
                ID : {product.id}
              </Typography>
              <Stack sx={{ ml: 2, pl: 2, borderLeft: '2px solid grey' }}>
                <Typography variant="body1" sx={{ color: 'inherit' }}>
                  {product.name}
                </Typography>
              </Stack>
            </Stack>
            <Stack direction="row">
              <Stack sx={{ flex: 0.3 }}>
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ color: 'inherit' }}>
                  Số lượng &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;: {product.quantity}
                </Typography>
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ color: 'green' }}>
                  Giá khởi điểm :{' '}
                  {product.min_price.toLocaleString('tr-TR', {
                    style: 'currency',
                    currency: 'VND',
                  })}
                </Typography>
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ color: 'red' }}>
                  Giá tối đa &nbsp; &nbsp; &nbsp; &nbsp; :{' '}
                  {product.expect_price.toLocaleString('tr-TR', {
                    style: 'currency',
                    currency: 'VND',
                  })}
                </Typography>
              </Stack>
              <Stack sx={{ flex: 0.7 }}>
                <Typography
                  fontSize={'0.9rem'}
                  variant="body2"
                  sx={{
                    color: 'inherit',
                    overflow: 'hidden',
                    textOverflow: 'ellipsis',
                    display: '-webkit-box',
                    WebkitLineClamp: '3',
                    WebkitBoxOrient: 'vertical',
                  }}
                >
                  Mô tả &nbsp; &nbsp; &nbsp;: {product.description}
                </Typography>
              </Stack>
            </Stack>
          </Stack>
          {/* Right button */}
          <Stack alignItems="flex-end" sx={{ width: '20%' }}>
            {/* Update and Delete Address */}
            <UpdateProductForm
              key={product.id}
              product={product}
              handleDelete={handleDelete}
              handleUpdateDefault={handleUpdateDefault}
            />
          </Stack>
        </Stack>
        {productsData.length - index !== 1 && <Divider sx={{ mb: 2 }} />}
      </Stack>
    ))
  ) : (
    <>Loading...</>
  );
}
