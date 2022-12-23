import { useState, useEffect, lazy, Suspense, useContext } from 'react';
import { Link as RouterLink, useNavigate, useLocation, useSearchParams } from 'react-router-dom';
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

import useMediaQuery from '@mui/material/useMediaQuery';
import { LoadingButton } from '@mui/lab';
import { styled, useTheme } from '@mui/material/styles';
import { Icon } from '@iconify/react';
import { FormProvider, RHFTextField } from '../../components/hook-form';
import useLocationForm from './useLocationForm';
import { ReloadContext } from '../../utils/Context';

const customStyles = {
  option: (provided, state) => ({
    ...provided,
    color: state.isSelected && '#F62217',
    backgroundColor: state.isSelected && 'white',
  }),
  control: (base, state) => ({
    ...base,
    borderRadius: 8,
    marginBottom: '4%',
    boxShadow: 'none',
    border: state.isFocused && '1px solid #F62217 !important',
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

export default function UpdateAddressForm({ address, handleDelete, handleUpdateDefault }) {
  const theme = useTheme();
  const navigate = useNavigate();

  const { isReloading, setIsReloading } = useContext(ReloadContext);

  // User information
  const [firstName, setFirstName] = useState(address.firstname);
  const [lastName, setLastName] = useState(address.lastname);
  const [phoneNumber, setPhoneNumber] = useState(address.phone);
  const [provinceData, setProvinceData] = useState(address.province);
  const [districtData, setDistrictData] = useState(address.district);
  const [subDistrictData, setSubDistrictData] = useState(address.sub_district);
  const [addressDetail, setAddressDetail] = useState(address.address);
  const [addressType, setAddressType] = useState(address.type_address);
  const [isDefault, setIsDefault] = useState(address.is_default);

  const stringToBoolean = (value) => {
    return String(value) === '1' || String(value).toLowerCase() === 'true';
  };
  const defaultValues = {
    firstname: '',
    lastname: '',
    phone: '',
    province: '',
    district: '',
    email: address.email,
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

  const onSubmit = async () => {
    const payload = {
      firstname: firstName,
      lastname: lastName,
      phone: phoneNumber,
      province: provinceData,
      district: districtData,
      email: address.email,
      sub_district: subDistrictData,
      address: addressDetail,
      type_address: addressType,
      is_default: stringToBoolean(isDefault),
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
        handleCloseUpdForm();
      }
      if (res.status === 401) {
        setError(true);
        setIsReloading(false);
        alert('You need to login first');
        navigate('/auctee/login', { replace: true });
      }
    });
  };
  const [error, setError] = useState(false);

  const fullScreen = useMediaQuery(theme.breakpoints.down('md'));

  const [openUpdForm, setOpenUpdForm] = useState(false);
  const [openDelForm, setOpenDelForm] = useState(false);

  // update form
  const handleClickOpenUpdForm = () => {
    setOpenUpdForm(true);
  };

  const handleCloseUpdForm = () => {
    setOpenUpdForm(false);
  };

  // delete form
  const handleClickOpenDelForm = () => {
    setOpenDelForm(true);
  };

  const handleCloseDelForm = () => {
    setOpenDelForm(false);
  };
  //  Location data
  const { state, onCitySelect, onDistrictSelect, onWardSelect } = useLocationForm(openUpdForm);
  const { cityOptions, districtOptions, wardOptions, selectedCity, selectedDistrict, selectedWard } = state;
  useEffect(() => {
    state.selectedCity = address.province;
    state.selectedDistrict = address.district;
    state.selectedWard = address.sub_district;
  }, [address, isReloading]);

  return (
    <>
      <Stack direction="row">
        <Button
          onClick={handleClickOpenUpdForm}
          sx={{
            borderRadius: 0,
            maxWidth: '80px',
            p: 0,
            bgcolor: 'transparent',
            border: 'none',
            fontSize: '0.9rem',
            color: 'green',
            fontWeight: 400,
            textTransform: 'none',
            '&:hover': {
              bgcolor: 'transparent',
              textDecoration: 'underline',
            },
          }}
        >
          Cập nhật
        </Button>
        <Dialog
          open={openUpdForm}
          sx={{ margin: 'auto', minWidth: '480px' }}
          BackdropProps={{
            style: { backgroundColor: 'rgba(0,0,30,0.15)' },
            invisible: true,
          }}
        >
          <FormProvider methods={methods} onSubmit={handleSubmit(onSubmit)}>
            <DialogTitle fontWeight={500}>Cập nhật địa chỉ</DialogTitle>
            <Stack sx={{ px: 3 }}>
              {/* Full name */}
              <Stack justifyContent="space-between" direction="row" sx={{ flex: 2, pb: 2 }}>
                <RHFTextField
                  color="error"
                  required
                  label="Tên"
                  name="firstname"
                  type="text"
                  value={firstName}
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
                  value={lastName}
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
                  value={phoneNumber}
                  onChange={(e) => setPhoneNumber(e.target.value)}
                  size="small"
                  variant="outlined"
                  sx={{}}
                />
              </Stack>
              {/* Province - District - Ward */}
              <Stack key={address.ID} sx={{ zIndex: 10000 }}>
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
                  placeholder={selectedCity || 'Tỉnh/Thành phố'}
                  value={selectedCity}
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
                  placeholder={selectedDistrict || 'Quận/Huyện'}
                  value={selectedDistrict}
                />

                <Select
                  styles={customStyles}
                  name="subdistrict"
                  key={`wardId_${selectedWard?.value}`}
                  isDisabled={wardOptions.length === 0}
                  options={wardOptions}
                  placeholder={selectedWard || 'Phường/Xã'}
                  onChange={(option) => {
                    onWardSelect(option);
                    setSubDistrictData(option.label);
                  }}
                  value={selectedWard}
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
                  value={addressDetail}
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
              {/* Address type */}
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
              {/* Default address */}
              <Stack>
                <FormControlLabel
                  disabled={address.is_default}
                  name="isdefault"
                  value={address.is_default}
                  control={<Checkbox size="small" onChange={(e) => setIsDefault(e.target.checked)} />}
                  label={<Typography variant="body2">Đặt làm địa chỉ mặc định</Typography>}
                />
              </Stack>
              {/* Buttons */}
              <Stack
                justifyContent="center"
                alignItems="center"
                direction="row"
                sx={{ mt: 2, pb: 4, position: 'relative' }}
              >
                <Button
                  size="medium"
                  variant="outlined"
                  onClick={handleCloseUpdForm}
                  sx={{
                    px: 1.6,
                    position: 'absolute',
                    right: 84,
                    color: 'inherit',
                    border: '1px solid white',
                    opacity: 0.85,
                    '&:hover': {
                      borderColor: 'black',
                      opacity: 1,
                    },
                  }}
                >
                  Trở lại
                </Button>
                <LoadingButton
                  key={address.ID}
                  disableRipple
                  color="error"
                  sx={{ px: 3, position: 'absolute', right: 1, bgcolor: '#F62217' }}
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
        {/* Delete Address */}
        {!address.is_default && (
          <Button
            key={address.ID}
            onClick={handleClickOpenDelForm}
            sx={{
              ml: 1,
              borderRadius: 0,
              p: 0,
              bgcolor: 'transparent',
              border: 'none',
              fontSize: '0.9rem',
              fontWeight: 400,
              '&:hover': {
                bgcolor: 'transparent',
                textDecoration: 'underline',
              },
            }}
          >
            Xoá
          </Button>
        )}
        {/* Dialog Delete */}
        <Dialog
          sx={{ margin: 'auto', minWidth: '480px' }}
          BackdropProps={{
            style: { backgroundColor: 'rgba(0,0,30,0.2)' },
            invisible: true,
          }}
          fullScreen={fullScreen}
          open={openDelForm}
        >
          <Stack sx={{ p: 3 }}>Bạn có chắc muốn xoá địa chỉ này?</Stack>

          <Stack sx={{ p: 2 }} justifyContent="flex-end" direction="row" alignItems="center">
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
              onClick={handleCloseDelForm}
            >
              Trở lại
            </Button>
            <Button
              key={address.ID}
              color="error"
              variant="contained"
              sx={{
                ml: 1,
                width: '62px',
                color: 'white',
                bgcolor: '#F62217',
              }}
              onClick={() => {
                handleDelete(address.ID);
                handleCloseDelForm();
              }}
              autoFocus
            >
              Xoá
            </Button>
          </Stack>
        </Dialog>
      </Stack>
      {/* Set default address button */}
      <Button
        key={address.ID}
        onClick={() => handleUpdateDefault(address)}
        disabled={address.is_default}
        variant="outlined"
        sx={{
          borderRadius: 0,
          my: 1,
          px: '10px !important',
          py: '2px !important',
          color: 'inherit',
          bgcolor: 'transparent',
          fontWeight: 400,
          textTransform: 'none',
          border: '1px solid black',
          '&:hover': {
            bgcolor: 'transparent',
            border: '1px solid grey',
            opacity: 0.85,
          },
        }}
        color="error"
        component="label"
      >
        Thiết lập mặc định
      </Button>
    </>
  );
}
