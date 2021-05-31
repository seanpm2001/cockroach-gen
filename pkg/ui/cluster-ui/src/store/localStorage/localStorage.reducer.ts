// Copyright 2018 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { DOMAIN_NAME } from "../utils";

export type LocalStorageState = {
  "adminUi/showDiagnosticsModal": boolean;
};

type Payload = {
  key: keyof LocalStorageState;
  value: any;
};

// TODO (koorosh): initial state should be restored from preserved keys in LocalStorage
const initialState: LocalStorageState = {
  "adminUi/showDiagnosticsModal": false,
};

const localStorageSlice = createSlice({
  name: `${DOMAIN_NAME}/localStorage`,
  initialState,
  reducers: {
    update: (state, action: PayloadAction<Payload>) => {
      state[action.payload.key] = action.payload.value;
    },
  },
});

export const { actions, reducer } = localStorageSlice;
