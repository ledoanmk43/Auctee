import { useState, useEffect } from 'react';
import { styled } from '@mui/material/styles';
import Countdown from 'react-countdown';
import { Typography, Stack } from '@mui/material';
import { red } from '@mui/material/colors';
import moment from 'moment';

const formatTime = (time) => {
  return String(time).padStart(2, '0');
};

const renderer = ({ days, hours, minutes, seconds, completed }) => {
  if (days === 1) {
    return (
      <span>
        {days} ngày {formatTime(hours)}:{formatTime(minutes)}:{formatTime(seconds)}
      </span>
    );
  }
  if (days !== 0) {
    return (
      <span>
        {days} ngày {formatTime(hours)}:{formatTime(minutes)}:{formatTime(seconds)}
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
const CountDown = ({ text, time }) => {
  const [remain, setRemain] = useState(time);
  const [done, setDone] = useState(false);
  useEffect(() => {
    setRemain(new Date(moment(time)).getTime());
    setDone(true);
  }, [time]);

  return (
    <>
      {done && (
        <Stack
          direction="row"
          sx={{
            whiteSpace: 'nowrap',
            maxWidth: `${text === 'Còn lại' ? '194px' : '100%'}`,

            justifyContent: 'space-between',
          }}
        >
          <Typography
            sx={{
              minWidth: '132px',
              height: '70%',
              color: 'red',
              fontSize: `${text === 'Còn lại:' ? '0.8rem' : 'inherit'}`,
            }}
            variant="caption"
          >
            {text} &nbsp;
            <Countdown renderer={renderer} date={remain} />
          </Typography>
        </Stack>
      )}
    </>
  );
};

export default CountDown;
