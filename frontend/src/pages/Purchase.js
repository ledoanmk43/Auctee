import { useState, useEffect, lazy } from 'react';
import { Link as RouterLink, useNavigate, useLocation, useSearchParams } from 'react-router-dom';
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
import { Icon } from '@iconify/react';
import { styled, useTheme } from '@mui/material/styles';
import { Box } from '@mui/system';
import account from '../API/account';
import { FormProvider, RHFTextField } from '../components/hook-form';

const Page = lazy(() => import('../components/Page'));

const RootStyle = styled('div')(({ theme }) => ({
  [theme.breakpoints.up('md')]: {
    display: 'flex',
    flexDirection: 'column',
    mx: '176x',
    backgroundColor: 'white',
    height: '100%',
    minHeight: '580px',
  },
}));
export default function Purchase() {
  const navigate = useNavigate();
  const location = useLocation();

  const [isFetching, setIsFetching] = useState(true);
  const [userData, setUserData] = useState();
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

  // Get user's data base on access_token
  const handleFetchUserData = async () => {
    await fetch('http://localhost:1001/auctee/user/profile', {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
    }).then((res) => {
      if (res.status === 200) {
        res.json().then((data) => {
          const today = new Date();
          setUserData(data);
          const date = Number(data.birthday.slice(0, 2));
          const month = Number(data.birthday.slice(3, 5));
          const year = data.birthday.slice(6, 10);
          setDate(date > 0 ? date : String(today.getDate()).padStart(2, '0'));
          setMonth(month > 0 ? month : String(today.getMonth() + 1).padStart(2, '0'));
          setYear(year > 0 ? year : today.getFullYear());
          setShopName(data.shopname);
          setNickName(data.nickname);
          setIsMale(data.gender);
          setPhoneNumber(data.phone);
          setAvatarFile(data.avatar);
          setIsFetching(false);
        });
      }
      if (res.status === 401) {
        alert('You need to login first');
        setIsFetching(true);
        navigate('/auctee/login', { replace: true });
      }
    });
  };

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

  const onSubmit = async () => {
    const payload = {};

    // await fetch('http://localhost:1001/auctee/user/profile/setting', {
    //   method: 'PUT',
    //   headers: { 'Content-Type': 'application/json' },
    //   body: JSON.stringify(payload),
    //   credentials: 'include',
    // }).then((res) => {
    //   if (res.status === 200) {
    //     setIsUpdated(true);
    //     navigate(0);
    //   }
    //   if (res.status === 409) {
    //     setIsUpdated(false);
    //   }
    //   if (res.status === 400) {
    //     setIsUpdated(false);
    //   }
    // });
  };

  useEffect(() => {
    handleFetchUserData();
  }, [isFetching]);

  return !isFetching ? (
    <Page title="Thanh toán">
      <RootStyle sx={{ px: 3, py: 2 }}>
        {/* Heading */}
        <Stack>
          <Typography fontSize={'1.2rem'} variant="body2" sx={{ color: 'black' }}>
            Đơn mua và Số dư
          </Typography>
          <Typography variant="body2" sx={{ color: 'black', opacity: 0.8, pb: 2 }}>
            Luôn duy trì số dư trong ví để thuận tiện mua sắm
          </Typography>
        </Stack>
        {/* Main */}
        <FormProvider methods={methods} onSubmit={handleSubmit(onSubmit)}>
          <Stack direction="row" sx={{ p: 2, boxShadow: 4, borderRadius: 2 }}>
            <Stack sx={{ flex: 2 }}>
              <Stack sx={{ ml: 0.5 }}>
                <Stack direction="row">
                  <Icon icon="bi:coin" color="#eba123" fontSize="3rem" />
                  <Stack sx={{ ml: 1 }}>
                    <Typography fontSize="1.2rem" color="#eba123">
                      {userData.total_income.toLocaleString('tr-TR', {
                        style: 'currency',
                        currency: 'VND',
                      })}
                    </Typography>
                    <Typography fontStyle="italic" variant="caption" fontSize="0.9rem" sx={{ opacity: 0.7 }}>
                      Số dư hiện tại
                    </Typography>
                  </Stack>
                </Stack>
                <Typography
                  fontStyle="italic"
                  variant="body2"
                  sx={{ color: '#f44336', minWidth: '100px', opacity: 0.9, mt: 2 }}
                >
                  Mẹo: &nbsp;Số dư trong ví phải lớn hơn giá trị của sản phầm bạn muốn tham gia đấu giá
                </Typography>
              </Stack>
            </Stack>
            <Stack justifyContent="center" alignItems="flex-start" direction="row">
              <LoadingButton
                disableRipple
                color="error"
                sx={{ px: 3, textTransform: 'none' }}
                size="medium"
                type="submit"
                variant="contained"
                loading={isSubmitting}
              >
                <Icon icon="bi:coin" /> &nbsp; Nạp tiền vào ví
              </LoadingButton>
            </Stack>
          </Stack>
        </FormProvider>
        
      </RootStyle>
    </Page>
  ) : (
    <>Có lỗi xảy ra</>
  );
}
