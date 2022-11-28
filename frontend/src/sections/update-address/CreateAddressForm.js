import { useState, useEffect, lazy, Suspense, useContext } from 'react';
import { Link as RouterLink, useNavigate, useLocation, useSearchParams, useOutletContext } from 'react-router-dom';
import { useForm } from 'react-hook-form';
import Select from 'react-select';

// material
import {
  Dialog,
  Checkbox,
  DialogTitle,
  RadioGroup,
  Radio,
  Typography,
  Stack,
  Button,
  Divider,
  Box,
  alertTitleClasses,
  FormControlLabel,
} from '@mui/material';
import { LoadingButton } from '@mui/lab';
import { styled, useTheme } from '@mui/material/styles';
import { Icon } from '@iconify/react';
import { FormProvider, RHFTextField } from '../../components/hook-form';
import useLocationForm from './useLocationForm';
import { ReloadContext } from '../../utils/Context';

const customStyles = {
  option: (provided, state) => ({
    ...provided,
    color: state.isSelected && '#f44336',
    backgroundColor: state.isSelected && 'white',
  }),
  control: (base, state) => ({
    ...base,
    borderRadius: 8,
    marginBottom: '4%',
    boxShadow: 'none',
    border: state.isFocused && '1px solid #f44336 !important',
    '&:hover': {
      border: '1px solid black',
    },
  }),
  menu: (base) => ({
    ...base,
    marginTop: '-4%',
  }),
  menuList: (base) => ({
    ...base,

    marginTop: 0,
    maxHeight: '200px',
  }),
};

export default function CreateAddressForm() {
  const theme = useTheme();
  const navigate = useNavigate();
  const { isReloading, setIsReloading } = useContext(ReloadContext);
  //  Location data
  const { state, onCitySelect, onDistrictSelect, onWardSelect } = useLocationForm(false);

  const { cityOptions, districtOptions, wardOptions, selectedCity, selectedDistrict, selectedWard } = state;

  // User information
  const [firstName, setFirstName] = useState('');
  const [lastName, setLastName] = useState('');
  const [phoneNumber, setPhoneNumber] = useState('');
  const [provinceData, setProvinceData] = useState('');
  const [districtData, setDistrictData] = useState('');
  const [subDistrictData, setSubDistrictData] = useState('');
  const [addressDetail, setAddressDetail] = useState('');
  const [addressType, setAddressType] = useState('');
  const [isDefault, setIsDefault] = useState(false);

  const [isFetching, setIsFetching] = useState(true);
  const userData = useOutletContext();

  const stringToBoolean = (value) => {
    return String(value) === '1' || String(value).toLowerCase() === 'true';
  };
  const defaultValues = {
    firstname: '',
    lastname: '',
    phone: '',
    province: '',
    district: '',
    subdistrict: '',
    addressdetail: '',
    addresstype: '',
    isdefault: isDefault,
  };

  const methods = useForm({
    defaultValues,
  });
  const {
    handleSubmit,
    formState: { isSubmitting },
  } = methods;

  const [isUpdated, setIsUpdated] = useState(false);
  const [error, setError] = useState(false);
  
  const onSubmit = async () => {
    const payload = {
      firstname: firstName,
      lastname: lastName,
      phone: phoneNumber,
      email: userData.email,
      province: provinceData,
      district: districtData,
      sub_district: subDistrictData,
      address: addressDetail,
      type_address: addressType,
      is_default: stringToBoolean(isDefault),
    };
    await fetch('http://localhost:1001/auctee/user/address', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
      credentials: 'include',
    }).then((res) => {
      if (res.status === 200) {
        setIsUpdated(true);
        setError(false);
        setIsReloading(true);
        setOpen(false);
      }
      if (res.status === 400 || res.status === 409) {
        setIsUpdated(false);
        setError(true);
      }
    });
  };

  const [open, setOpen] = useState(false);

  const handleClickOpen = () => {
    setOpen(true);
  };
  useEffect(() => {
    setIsReloading(false);
  }, [isFetching, open, isReloading]);

  return (
    <>
      <Stack justifyContent="space-between" alignItems="center" direction="row" sx={{ maxHeight: '100%', pb: 2 }}>
        <Typography fontSize={'1.2rem'} variant="body2" sx={{ color: 'black' }}>
          Địa chỉ
        </Typography>
        <Button
          onClick={handleClickOpen}
          sx={{
            my: 1,
            px: '20px !important',
            color: 'white',
            bgcolor: theme.palette.background.main,
            fontWeight: 500,
            textTransform: 'none',
          }}
          color="error"
          variant="contained"
          component="label"
        >
          <Icon icon="akar-icons:plus" />
          &nbsp; Thêm địa chỉ mới
        </Button>
      </Stack>
      <Dialog
        open={open}
        sx={{ margin: 'auto', minWidth: '480px' }}
        BackdropProps={{
          style: { backgroundColor: 'rgba(0,0,30,0.4)' },
          invisible: true,
        }}
      >
        <DialogTitle fontWeight={500}>Địa chỉ mới</DialogTitle>
        <FormProvider methods={methods} onSubmit={handleSubmit(onSubmit)}>
          <Stack sx={{ px: 3 }}>
            {/* name */}
            <Stack justifyContent="space-between" direction="row" sx={{ flex: 2, pb: 2 }}>
              <RHFTextField
                color="error"
                required
                label="Tên"
                name="firstname"
                type="text"
                value={firstName || ''}
                onChange={(e) => setFirstName(e.target.value)}
                size="small"
                variant="outlined"
                sx={{ width: '30%', borderRadius: 1 }}
              />
              <RHFTextField
                color="error"
                required
                label="Họ"
                name="lastname"
                type="text"
                value={lastName || ''}
                onChange={(e) => setLastName(e.target.value)}
                size="small"
                variant="outlined"
                sx={{ width: '60%' }}
              />
            </Stack>
            {/* Phone */}
            <Stack direction="row" sx={{ flex: 2, pb: 2 }}>
              <RHFTextField
                color="error"
                required
                label="Số điện thoại"
                name="phone"
                type="number"
                value={phoneNumber || ''}
                onChange={(e) => setPhoneNumber(e.target.value)}
                size="small"
                variant="outlined"
                sx={{}}
              />
            </Stack>
            {/* Province */}
            <Stack sx={{ zIndex: 10000 }}>
              <Select
                styles={customStyles}
                name="province"
                key={`cityId_${selectedCity?.value}`}
                isDisabled={cityOptions.length === 0}
                options={cityOptions}
                onChange={(option) => {
                  onCitySelect(option);
                  setProvinceData(option.label);
                }}
                placeholder="Tỉnh/Thành"
                defaultValue={selectedCity}
              />

              <Select
                styles={customStyles}
                name="district"
                key={`districtId_${selectedDistrict?.value}`}
                isDisabled={districtOptions.length === 0}
                options={districtOptions}
                onChange={(option) => {
                  onDistrictSelect(option);
                  setDistrictData(option.label);
                }}
                placeholder="Quận/Huyện"
                defaultValue={selectedDistrict}
              />

              <Select
                styles={customStyles}
                name="subdistrict"
                key={`wardId_${selectedWard?.value}`}
                isDisabled={wardOptions.length === 0}
                options={wardOptions}
                placeholder="Phường/Xã"
                onChange={(option) => {
                  onWardSelect(option);
                  setSubDistrictData(option.label);
                }}
                defaultValue={selectedWard}
              />
            </Stack>
            {/* Address Detail */}
            <Stack direction="row" sx={{ flex: 2, pb: 2 }}>
              <RHFTextField
                color="error"
                required
                label="Địa chỉ cụ thể"
                name="addressdetail"
                type="text"
                value={addressDetail || ''}
                onChange={(e) => setAddressDetail(e.target.value)}
                size="small"
                variant="outlined"
              />
            </Stack>
            <Stack alignItems="center" direction="row" sx={{ height: 3 }}>
              {error && (
                <Typography variant="body2" color="error">
                  Thông tin địa chỉ không được để trống
                </Typography>
              )}
            </Stack>
            <Stack>
              <Typography fontSize="0.9rem" sx={{ mt: 1 }}>
                Loại địa chỉ:
              </Typography>
              <Stack direction="row">
                <RadioGroup required name="addresstype" value={addressType} row sx={{}}>
                  <FormControlLabel
                    value="nhà riêng"
                    onChange={(e) => setAddressType(e.target.value)}
                    control={<Radio size="small" />}
                    label="Nhà riêng"
                  />
                  <FormControlLabel
                    value="văn phòng"
                    onChange={(e) => setAddressType(e.target.value)}
                    control={<Radio size="small" />}
                    label="Văn phòng"
                  />
                </RadioGroup>
              </Stack>
            </Stack>
            <Stack>
              <FormControlLabel
                name="isdefault"
                control={<Checkbox size="small" onChange={(e) => setIsDefault(e.target.checked)} />}
                label={<Typography variant="body2">Đặt làm địa chỉ mặc định</Typography>}
              />
            </Stack>
            <Stack
              justifyContent="center"
              alignItems="center"
              direction="row"
              sx={{ mt: 2, pb: 4, position: 'relative' }}
            >
              <Button
                size="medium"
                variant="outlined"
                onClick={() => setOpen(false)}
                sx={{
                  px: 1.6,
                  position: 'absolute',
                  right: 84,
                  color: 'inherit',
                  border: '1px solid white',
                  opacity: 0.85,
                  textTransform: 'none',
                  '&:hover': {
                    borderColor: 'black',
                    opacity: 1,
                  },
                }}
              >
                Trở lại
              </Button>
              <LoadingButton
                disableRipple
                color="error"
                sx={{ px: 3, position: 'absolute', right: 1 }}
                size="medium"
                type="submit"
                variant="contained"
                loading={isSubmitting}
              >
                Lưu
              </LoadingButton>
            </Stack>
          </Stack>
        </FormProvider>
      </Dialog>
    </>
  );
}
