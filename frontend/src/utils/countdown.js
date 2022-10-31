import { useState, useEffect } from 'react';
import { styled } from '@mui/material/styles';
import Countdown from 'react-countdown';
import { Typography, Stack } from '@mui/material';
import { red } from '@mui/material/colors';
import moment from 'moment';
import Iconify from '../components/Iconify';

const formatTime = (time) => {
  return String(time).padStart(2, '0');
};

const Keyframes = styled('div')({
  '@keyframes pull': {
    '0%': {
      transform: 'rotate(-70deg)',
    },
    '25%': {
      transform: 'rotate(-40deg)',
    },
    '50%': {
      transform: 'rotate(0deg)',
    },
    '75%': {
      transform: 'rotate(40deg)',
    },
    '100%': {
      transform: 'rotate(-70deg)',
    },
  },
  animation: 'pull 1s infinite ease-in-out',
  position: 'absolute',
  top: '16px',
  right: '-18px',
  transform: 'rotate(45deg)',
});

const renderer = ({ days, hours, minutes, seconds, completed }) => {
  if (days === 1) {
    return (
      <span>
        {days} day {formatTime(hours)}:{formatTime(minutes)}:{formatTime(seconds)}
      </span>
    );
  }
  if (days !== 0) {
    return (
      <span>
        {days} days {formatTime(hours)}:{formatTime(minutes)}:{formatTime(seconds)}
      </span>
    );
  }
  if (days !== 1 || days !== 0) {
    return (
      <span>
        {formatTime(hours)}:{formatTime(minutes)}:{formatTime(seconds)}
      </span>
    );
  }
};
const CountDown = ({ time }) => {
  const [remain, setRemain] = useState(time);
  const [done, setDone] = useState(false);
  useEffect(() => {
    setRemain(new Date(moment(time)).getTime());
    setDone(true);
  }, [time]);

  return (
    <>
      {done && (
        <Stack direction="row" sx={{ maxWidth: 194, position: 'relative', justifyContent: 'space-between' }}>
          <Typography sx={{ minWidth: '132px', height: '70%' }} variant="caption" color="red">
            End in: &nbsp;
            <Countdown renderer={renderer} date={remain} />
          </Typography>
          <Stack color="brown" sx={{ position: 'relative' }} direction="row">
            <Keyframes>
              <Iconify sx={{ transform: 'rotate(30deg)' }} icon={'mingcute:auction-line'} width={25} height={25} />
            </Keyframes>
          </Stack>
        </Stack>
      )}
    </>
  );
};

export default CountDown;
