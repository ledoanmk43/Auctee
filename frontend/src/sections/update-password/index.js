import { useState, useEffect, lazy } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';
import { useForm } from 'react-hook-form';
import { LoadingButton } from '@mui/lab';
// material
import { Typography, Stack, IconButton, InputAdornment } from '@mui/material';
import * as Yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { styled, useTheme } from '@mui/material/styles';
import Iconify from '../../components/Iconify';
import { FormProvider, RHFTextField } from '../../components/hook-form';

export default function ChangePasswordForm() {
  const navigate = useNavigate();

  const [showOldPassword, setShowOldPassword] = useState(false);
  const [showNewPassword, setShowNewPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);

  const [isUpdated, setIsUpdated] = useState(false);
  const [isWrongPwd, setIsWrongPwd] = useState(false);
  const [isSamePwd, setIsSamePwd] = useState(false);
  const [isWrongConfirm, setIsWrongConfirm] = useState(false);

  const ChangePwdSchema = Yup.object().shape({
    // oldpassword: Yup.string().required('required'),
    // newpassword: Yup.string().min(6).max(32).required('required'),
    // confirmpassword: Yup.string().required('required'),
  });
  const defaultValues = {
    oldpassword: '',
    newpassword: '',
    confirmpassword: '',
  };

  const methods = useForm({
    resolver: yupResolver(ChangePwdSchema),
    defaultValues,
  });
  const {
    handleSubmit,
    formState: { isSubmitting },
  } = methods;

  const onSubmit = async (data) => {
    if (data.confirmpassword !== data.newpassword) {
      setIsWrongConfirm(true);
      return;
    }
    const payload = {
      old_password: data.oldpassword,
      new_password: data.newpassword,
    };

    await fetch('http://localhost:8080/auctee/user/profile', {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
      credentials: 'include',

      mode: 'cors',
    }).then((res) => {
      if (res.status === 200) {
        setIsUpdated(true);
        setIsWrongPwd(false);
        setIsSamePwd(false);
        setIsWrongConfirm(false);
        navigate(0);
      }
      if (res.status === 409) {
        setIsSamePwd(true);
        setIsUpdated(false);
        setIsWrongPwd(false);
        setIsWrongConfirm(false);
      }
      if (res.status === 400) {
        setIsWrongPwd(true);
        setIsSamePwd(false);
        setIsUpdated(false);
        setIsWrongConfirm(false);
      }
    });
  };
  useEffect(() => {}, [isUpdated]);
  return (
    <FormProvider
      methods={methods}
      onSubmit={handleSubmit(onSubmit)}
      direction="row"
      sx={{ width: '50%', justifyContent: 'center' }}
    >
      {/* Old password */}
      <Stack alignItems="center" direction="row" sx={{ maxWidth: '570px', pb: 3, minHeight: '64px', mt: 2 }}>
        <Typography textAlign="right" variant="body2" sx={{ color: 'black', opacity: 0.8, minWidth: '200px' }}>
          Mật khẩu cũ
        </Typography>
        <RHFTextField
          required
          name="oldpassword"
          type={showOldPassword ? 'text' : 'password'}
          size="small"
          variant="outlined"
          sx={{
            ml: 3,
            width: '200px',
            px: 2,
          }}
          InputProps={{
            endAdornment: (
              <InputAdornment position="end">
                <IconButton edge="end" onClick={() => setShowOldPassword(!showOldPassword)}>
                  <Iconify sx={{ fontSize: '1.1rem' }} icon={showOldPassword ? 'eva:eye-fill' : 'eva:eye-off-fill'} />
                </IconButton>
              </InputAdornment>
            ),
          }}
        />
        {isWrongPwd && (
          <Typography
            textAlign="
                left"
            variant="caption"
            sx={{ color: 'red', minWidth: '100px' }}
          >
            * Sai mật khẩu
          </Typography>
        )}
      </Stack>
      {/* New password */}
      <Stack alignItems="center" direction="row" sx={{ maxWidth: '570px', pb: 3 }}>
        <Typography textAlign="right" variant="body2" sx={{ color: 'black', opacity: 0.8, minWidth: '200px' }}>
          Mật khẩu mới
        </Typography>
        <RHFTextField
          required
          name="newpassword"
          type={showNewPassword ? 'text' : 'password'}
          size="small"
          variant="outlined"
          sx={{
            ml: 3,
            minWidth: '200px !important',
            maxWidth: '200px !important',
            px: 2,
          }}
          InputProps={{
            endAdornment: (
              <InputAdornment position="end">
                <IconButton edge="end" onClick={() => setShowNewPassword(!showNewPassword)}>
                  <Iconify sx={{ fontSize: '1.1rem' }} icon={showNewPassword ? 'eva:eye-fill' : 'eva:eye-off-fill'} />
                </IconButton>
              </InputAdornment>
            ),
          }}
        />
        {isSamePwd && (
          <Typography
            textAlign="
                left"
            variant="caption"
            sx={{ color: 'red', minWidth: '150px' }}
          >
            * Mật khẩu mới không được trùng mật khẩu cũ
          </Typography>
        )}
      </Stack>
      {/* Confirm new password */}
      <Stack alignItems="center" direction="row" sx={{ maxWidth: '570px', pb: 3 }}>
        <Typography textAlign="right" variant="body2" sx={{ color: 'black', opacity: 0.8, minWidth: '200px' }}>
          Xác nhận mật khẩu mới
        </Typography>
        <RHFTextField
          required
          name="confirmpassword"
          type={showConfirmPassword ? 'text' : 'password'}
          size="small"
          variant="outlined"
          sx={{
            ml: 3,
            width: '200px !important',
            px: 2,
          }}
          InputProps={{
            endAdornment: (
              <InputAdornment position="end">
                <IconButton edge="end" onClick={() => setShowConfirmPassword(!showConfirmPassword)}>
                  <Iconify
                    sx={{ fontSize: '1.1rem' }}
                    icon={showConfirmPassword ? 'eva:eye-fill' : 'eva:eye-off-fill'}
                  />
                </IconButton>
              </InputAdornment>
            ),
          }}
        />
        {isWrongConfirm && (
          <Typography
            textAlign="
                left"
            variant="caption"
            sx={{ color: 'red', minWidth: '150px' }}
          >
            * Mật khẩu không khớp
          </Typography>
        )}
      </Stack>
      {isUpdated && (
        <Stack alignItems="right" direction="row" sx={{ maxWidth: '570px', pb: 3, ml: 30 }}>
          <Typography
            textAlign="
                left"
            variant="caption"
            sx={{ minWidth: '150px', color: 'green' }}
          >
            Thay đổi mật khẩu thành công
          </Typography>
        </Stack>
      )}
      <Stack justifyContent="center" alignItems="center" direction="row" sx={{ pb: 4, mr: 55 }}>
        <LoadingButton
          disableRipple
          color="error"
          sx={{ px: 3 }}
          size="medium"
          type="submit"
          variant="contained"
          loading={isSubmitting}
        >
          Lưu
        </LoadingButton>
      </Stack>
    </FormProvider>
  );
}
