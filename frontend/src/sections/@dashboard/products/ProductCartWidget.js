import { useState, useEffect, lazy } from 'react';
// material
import { useTheme } from '@mui/material/styles';

import { Stack, Typography } from '@mui/material';

const ProductBadge = lazy(() => import('./ProductBadge'));
// component

// ----------------------------------------------------------------------

export default function CartWidget() {
  const theme = useTheme();

  const [activeAuctions, setActiveAuctions] = useState([]);
  const [loaded, setLoaded] = useState(false);
  const fetchCurrentAuctions = async () => {
    await fetch(`http://localhost:1009/auctee/user/all-current-bids`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
    }).then((res) => {
      if (res.status === 200) {
        res.json().then((data) => {
          setActiveAuctions(data);
          setLoaded(true);
        });
      }
    });
  };

  useEffect(() => {
    // eslint-disable-next-line no-unused-expressions
    !loaded && fetchCurrentAuctions();
  }, [activeAuctions, setLoaded]);

  return (
    <Stack
      sx={{
        maxHeight: '555px',
        overflow: 'auto',
        scrollbarWidth: 'thin',
        '&::-webkit-scrollbar': {
          width: '0.4em',
        },
        '&::-webkit-scrollbar-track': {
          background: '#f0e7e6',
        },
        '&::-webkit-scrollbar-thumb': {
          backgroundColor: '#cfc9c8',
        },
        '&::-webkit-scrollbar-thumb:hover': {
          background: '#bab3b1',
        },
        zIndex: 999,
        right: 5,
        display: 'flex',
        flexDirection: 'column',
        cursor: 'pointer',
        position: 'fixed',
        alignItems: 'center',
        top: theme.spacing(14.5),
        paddingRight: theme.spacing(1),
      }}
    >
      <Typography variant="button" sx={{ mt: '-5px', opacity: 0.75, textTransform: 'none' }}>
        Đang tham gia
      </Typography>
      {activeAuctions &&
        activeAuctions.map((auction, index) => <ProductBadge setLoaded={setLoaded} key={index} auction={auction} />)}
    </Stack>
  );
}
