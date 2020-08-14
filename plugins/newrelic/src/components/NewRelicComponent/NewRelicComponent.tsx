/*
 * Copyright 2020 Spotify AB
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import React, { FC } from 'react';
import { Grid } from '@material-ui/core';
import {
  Header,
  Page,
  pageTheme,
  Content,
  ContentHeader,
  HeaderLabel,
  SupportButton,
} from '@backstage/core';
import NewRelicFetchComponent from '../NewRelicFetchComponent';

const NewRelicComponent: FC<{}> = () => (
  <Page theme={pageTheme.tool}>
    <Header title="New Relic">
      <HeaderLabel label="Owner" value="Engineering" />
    </Header>
    <Content>
      <ContentHeader title="New Relic">
        <SupportButton>
          New Relic Application Performance Monitoring
        </SupportButton>
      </ContentHeader>
      <Grid container spacing={3} direction="column">
        <Grid item>
          <NewRelicFetchComponent />
        </Grid>
      </Grid>
    </Content>
  </Page>
);

export default NewRelicComponent;
