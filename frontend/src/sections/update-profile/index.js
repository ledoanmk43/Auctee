import { useState, useEffect, lazy } from 'react';
import { Link as RouterLink, useNavigate, useLocation, useSearchParams, useOutletContext } from 'react-router-dom';
import { useForm } from 'react-hook-form';
import FileBase64 from 'react-file-base64';
// material
import {
  Container,
  Avatar,
  FormGroup,
  TextField,
  RadioGroup,
  Radio,
  FormControlLabel,
  Typography,
  Stack,
  Button,
  Select,
  MenuItem,
  Divider,
} from '@mui/material';
import { LoadingButton } from '@mui/lab';

import { styled, useTheme } from '@mui/material/styles';
import { Box } from '@mui/system';
import account from '../../API/account';
import { FormProvider, RHFTextField } from '../../components/hook-form';

export default function UpdateProfileForm() {
  const navigate = useNavigate();
  const location = useLocation();
  const userData = useOutletContext();

  const [isFetching, setIsFetching] = useState(true);
  const [shopName, setShopName] = useState('');
  const [nickName, setNickName] = useState('');
  const [phoneNumber, setPhoneNumber] = useState('');
  const [avatarFile, setAvatarFile] = useState();
  const [HonorPoint, setHonorPoint] = useState(0);

  const [isMale, setIsMale] = useState(false); // 1 male : 0 female
  const dates = [
    1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
  ];
  const months = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12];
  const generateArrayOfYears = () => {
    const max = new Date().getFullYear();
    const min = max - 100;
    const years = [];

    for (let i = max; i >= min; i -= 1) {
      years.push(i);
    }
    return years;
  };
  const years = generateArrayOfYears();

  const [date, setDate] = useState('');
  const [month, setMonth] = useState('');
  const [year, setYear] = useState('');

  const [birthday, setBirthDay] = useState('');

  const [isUpdated, setIsUpdated] = useState(false);

  const defaultValues = {
    nickname: '',
    shopname: '',
    gender: isMale,
    phone: '',
    date: '',
    month: '',
    year: '',
    avatar: '',
  };

  const methods = useForm({
    defaultValues,
  });
  const {
    handleSubmit,
    formState: { isSubmitting },
  } = methods;

  const stringToBoolean = (value) => {
    return String(value) === '1' || String(value).toLowerCase() === 'true';
  };

  const onSubmit = async () => {
    const payload = {
      nickname: nickName,
      shopname: shopName,
      phone: phoneNumber,
      gender: stringToBoolean(isMale),
      birthday: `${date.length > 2 ? date : String(date).padStart(2, '0')}/${
        month.length > 2 ? month : String(month).padStart(2, '0')
      }/${year}`,
      avatar: avatarFile,
    };

    await fetch('http://localhost:1001/auctee/user/profile/setting', {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
      credentials: 'include',
    }).then((res) => {
      if (res.status === 200) {
        setIsUpdated(true);
        navigate(0);
      }
      if (res.status === 409) {
        setIsUpdated(false);
      }
      if (res.status === 400) {
        setIsUpdated(false);
      }
    });
  };

  const handleUpdateAvatar = (event) => {
    const file = event.target.files[0];
    if (file.size > 2000000) {
      alert('file too large');
      return;
    }
    const reader = new FileReader();
    reader.readAsDataURL(file);

    reader.onload = () => {
      setAvatarFile(reader.result); // base64encoded string
    };
    reader.onerror = (error) => {
      console.log('Error: ', error);
    };
  };

  useEffect(() => {
    if (userData) {
      const today = new Date();
      const dateConverted = Number(userData.birthday.slice(0, 2));
      const monthConverted = Number(userData.birthday.slice(3, 5));
      const yearConverted = userData.birthday.slice(6, 10);
      setDate(dateConverted > 0 ? dateConverted : String(today.getDate()).padStart(2, '0'));
      setMonth(monthConverted > 0 ? monthConverted : String(today.getMonth() + 1).padStart(2, '0'));
      setYear(yearConverted > 0 ? yearConverted : today.getFullYear());
      setShopName(userData.shopname);
      setNickName(userData.nickname);
      setIsMale(userData.gender);
      setPhoneNumber(userData.phone);
      setAvatarFile(userData.avatar);
      setIsFetching(false);
    }
  }, [userData]);

  return userData ? (
    <FormProvider methods={methods} onSubmit={handleSubmit(onSubmit)}>
      <Stack direction="row" sx={{ p: 2 }}>
        <Stack sx={{ flex: 2 }}>
          <Stack direction="row" fontStyle="italic" sx={{ pb: 3, ml: 0.5 }}>
            <Typography variant="body2" sx={{ color: '#f44336', minWidth: '100px', opacity: 0.9 }}>
              Mẹo: &nbsp;Duy trì điểm uy tín trên 80 để được tham gia đấu giá
            </Typography>
          </Stack>
          {/* Username */}
          <Stack direction="row" sx={{ pb: 3 }}>
            <Typography textAlign="right" variant="body2" sx={{ color: 'black', opacity: 0.8, minWidth: '100px' }}>
              Tên đăng nhập
            </Typography>
            <Typography variant="body2" sx={{ color: 'black', pl: 5 }}>
              {userData.username || ''}
            </Typography>
          </Stack>
          {/* Nickname */}
          <Stack alignItems="center" direction="row" sx={{ pb: 3 }}>
            <Typography textAlign="right" variant="body2" sx={{ color: 'black', opacity: 0.8, minWidth: '100px' }}>
              Tên
            </Typography>
            <RHFTextField
              required
              name="nickname"
              type="text"
              value={nickName || ''}
              onChange={(e) => setNickName(e.target.value)}
              size="small"
              variant="outlined"
              sx={{
                ml: 3,
                minWidth: '60%',
                px: 2,
              }}
            />
          </Stack>
          {/* Email */}
          <Stack direction="row" sx={{ pb: 3 }}>
            <Typography textAlign="right" variant="body2" sx={{ color: 'black', opacity: 0.8, minWidth: '100px' }}>
              Email
            </Typography>
            <Typography variant="body2" sx={{ color: 'black', pl: 5 }}>
              {String(userData.email).slice(-12).padStart(String(userData.email).length, '*')}
            </Typography>
          </Stack>
          {/* Phone number */}
          <Stack direction="row" sx={{ pb: 3 }}>
            <Typography textAlign="right" variant="body2" sx={{ color: 'black', opacity: 0.8, minWidth: '100px' }}>
              Số điện thoại
            </Typography>
            {userData.phone.length > 0 ? (
              <Typography variant="body2" sx={{ color: 'black', pl: 5 }}>
                {String(userData.phone).slice(-2).padStart(String(userData.phone).length, '*')}
              </Typography>
            ) : (
              <RHFTextField
                name="phone"
                required
                value={phoneNumber || ''}
                onChange={(e) => setPhoneNumber(e.target.value)}
                size="small"
                variant="outlined"
                sx={{
                  ml: 3,
                  minWidth: '60%',
                  px: 2,
                }}
              />
            )}
          </Stack>
          {/* Shop Name */}
          <Stack alignItems="center" direction="row" sx={{ pb: 3 }}>
            <Typography textAlign="right" variant="body2" sx={{ color: 'black', opacity: 0.8, minWidth: '100px' }}>
              Tên Shop
            </Typography>
            <RHFTextField
              name="shopname"
              required
              value={shopName || ''}
              onChange={(e) => setShopName(e.target.value)}
              size="small"
              variant="outlined"
              sx={{
                ml: 3,
                minWidth: '60%',
                px: 2,
              }}
            />
          </Stack>
          {/* Gender */}
          <Stack direction="row" alignItems="center" sx={{ pb: 3 }}>
            <Typography textAlign="right" variant="body2" sx={{ color: 'black', opacity: 0.8, minWidth: '100px' }}>
              Giới tính
            </Typography>
            <RadioGroup name="gender" value={isMale} row sx={{ ml: 5 }}>
              <FormControlLabel
                value="true"
                onChange={(e) => setIsMale(e.target.value)}
                control={<Radio size="small" />}
                label="Nam"
              />
              <FormControlLabel
                value="false"
                onChange={(e) => setIsMale(e.target.value)}
                control={<Radio size="small" />}
                label="Nữ"
              />
            </RadioGroup>
          </Stack>
          {/* Birthday */}
          <Stack direction="row" alignItems="center" sx={{ pb: 3 }}>
            <Typography textAlign="right" variant="body2" sx={{ color: 'black', opacity: 0.8, minWidth: '100px' }}>
              Ngày sinh
            </Typography>
            <Stack direction="row" sx={{ ml: 5 }}>
              {/* Date */}
              <Select
                name="date"
                value={date}
                onChange={(e) => setDate(e.target.value)}
                sx={{ maxHeight: '30px', mr: 2, minWidth: '100px' }}
              >
                {dates.map((item, index) => (
                  <MenuItem key={index} value={item}>
                    {item}
                  </MenuItem>
                ))}
              </Select>
              {/* Month */}
              <Select
                name="month"
                value={month}
                onChange={(e) => setMonth(e.target.value)}
                sx={{ maxHeight: '30px', mr: 2, minWidth: '100px' }}
              >
                {months.map((item, index) => (
                  <MenuItem key={index} value={item}>
                    Tháng {item}
                  </MenuItem>
                ))}
              </Select>
              {/* Year */}
              <Select
                name="year"
                value={year}
                onChange={(e) => setYear(e.target.value)}
                sx={{ maxHeight: '30px', mr: 2, minWidth: '100px' }}
              >
                {years.map((item, index) => (
                  <MenuItem key={index} value={item}>
                    {item}
                  </MenuItem>
                ))}
              </Select>
            </Stack>
          </Stack>
        </Stack>
        {/* Images */}
        <Stack justifyContent="space-around" direction="row" sx={{ flex: 1.5 }}>
          <Divider orientation="vertical" />
          <Stack alignItems="center" direction="column">
            <Avatar alt="Remy Sharp" src={avatarFile} sx={{ width: 120, height: 120 }} />
            <Button sx={{ my: 3, px: '20px !important' }} variant="outlined" color="error" component="label">
              Chọn Ảnh
              <input onChange={(e) => handleUpdateAvatar(e)} hidden accept="image/*" multiple type="file" />
            </Button>
            <Stack>
              <Typography variant="caption" sx={{ color: 'black' }}>
                Dụng lượng file tối đa 1 MB,
              </Typography>
              <Typography variant="caption" sx={{ color: 'black' }}>
                Định dạng:.JPEG, .PNG
              </Typography>
            </Stack>
          </Stack>
        </Stack>
      </Stack>
      <Stack justifyContent="center" alignItems="center" direction="row" sx={{ pb: 4, position: 'relative' }}>
        <LoadingButton
          disableRipple
          color="error"
          sx={{ px: 3, position: 'absolute', left: '156px' }}
          size="medium"
          type="submit"
          variant="contained"
          loading={isSubmitting}
        >
          Lưu
        </LoadingButton>
      </Stack>
    </FormProvider>
  ) : (
    <>Có lỗi xảy ra</>
  );
}
