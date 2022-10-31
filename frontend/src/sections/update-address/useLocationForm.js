import axios from 'axios';
import { useEffect, useState } from 'react';

const PATHS = {
  CITIES: 'https://raw.githubusercontent.com/nhidh99/codergamo/master/004-location-selects/locations/cities.json',
  DISTRICTS: 'https://raw.githubusercontent.com/nhidh99/codergamo/master/004-location-selects/locations/districts',
  WARDS: 'https://raw.githubusercontent.com/nhidh99/codergamo/master/004-location-selects/locations/wards',
  LOCATION: 'https://raw.githubusercontent.com/nhidh99/codergamo/master/004-location-selects/locations/location.json',
};

const FETCH_TYPES = {
  CITIES: 'FETCH_CITIES',
  DISTRICTS: 'FETCH_DISTRICTS',
  WARDS: 'FETCH_WARDS',
};

async function fetchLocationOptions(fetchType, locationId) {
  let url;
  switch (fetchType) {
    case FETCH_TYPES.CITIES: {
      url = PATHS.CITIES;
      break;
    }
    case FETCH_TYPES.DISTRICTS: {
      url = `${PATHS.DISTRICTS}/${locationId}.json`;
      break;
    }
    case FETCH_TYPES.WARDS: {
      url = `${PATHS.WARDS}/${locationId}.json`;
      break;
    }
    default: {
      return [];
    }
  }
  const res = await fetch(url);
  const locations = await res.json();

  return locations.data.map(({ id, name }) => ({ value: id, label: name }));
}

async function fetchInitialData(state) {
  let cityId;
  let districtId;
  await fetchLocationOptions(FETCH_TYPES.CITIES).then((res) => {
    cityId = res.find((c) => c.label === state.selectedCity).value;
  });

  await fetchLocationOptions(FETCH_TYPES.DISTRICTS, cityId).then((res) => {
    districtId = res.find((c) => c.label === state.selectedDistrict).value;
  });

  const [cities, districts, wards] = await Promise.all([
    fetchLocationOptions(FETCH_TYPES.CITIES),
    fetchLocationOptions(FETCH_TYPES.DISTRICTS, cityId),
    fetchLocationOptions(FETCH_TYPES.WARDS, districtId),
  ]);

  return {
    cityOptions: cities,
    districtOptions: districts,
    wardOptions: wards,
    selectedCity: state.selectedCity,
    selectedDistrict: state.selectedDistrict,
    selectedWard: state.selectedWard,
  };
}

function useLocationForm(shouldFetchInitialLocation) {
  const [state, setState] = useState({
    cityOptions: [],
    districtOptions: [],
    wardOptions: [],
    selectedCity: null,
    selectedDistrict: null,
    selectedWard: null,
  });

  const { selectedCity, selectedDistrict } = state;

  useEffect(() => {
    (async function () {
      if (shouldFetchInitialLocation) {
        await fetchInitialData(state).then((data) => {
          setState(data);
        });
      } else {
        const options = await fetchLocationOptions(FETCH_TYPES.CITIES);
        setState({ ...state, cityOptions: options });
      }
    })();
  }, [shouldFetchInitialLocation]);

  useEffect(() => {
    (async function () {
      if (!selectedCity) return;
      const options = await fetchLocationOptions(FETCH_TYPES.DISTRICTS, selectedCity.value);
      setState({ ...state, districtOptions: options });
    })();
  }, [selectedCity]);

  useEffect(() => {
    (async function () {
      if (!selectedDistrict) return;
      const options = await fetchLocationOptions(FETCH_TYPES.WARDS, selectedDistrict.value);
      setState({ ...state, wardOptions: options });
    })();
  }, [selectedDistrict]);

  function onCitySelect(option) {
    if (option !== selectedCity) {
      setState({
        ...state,
        districtOptions: [],
        wardOptions: [],
        selectedCity: option,
        selectedDistrict: null,
        selectedWard: null,
      });
    }
  }

  function onDistrictSelect(option) {
    if (option !== selectedDistrict) {
      setState({
        ...state,
        wardOptions: [],
        selectedDistrict: option,
        selectedWard: null,
      });
    }
  }

  function onWardSelect(option) {
    setState({ ...state, selectedWard: option });
  }

  return { state, onCitySelect, onDistrictSelect, onWardSelect };
}

export default useLocationForm;
