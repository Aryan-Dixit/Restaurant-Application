import styled from 'styled-components';

export const Row = styled.div`
  display: flex;
  flex-wrap: wrap;
  position: relative;
  margin-left: -15px;
  margin-right: -15px;
`;

export const Wrapper = styled.div`
  width: 100%;
  position: relative;
  margin-top: 20px;
`;

export const AgroTable = styled.table`
  font-size: 1rem;
  border-spacing: 0;
  width: 100%;
  border-radius: 2px;
  border-collapse: collapse;
  table-layout: fixed;
  border: 1px solid #000;
  > tbody > tr {
    border: 1px solid #000;
  }
  > thead > tr > th {
    height: 50px;
  }
  > thead > tr > th,
  > tbody > tr > td {
    padding: 10px;
    text-align: left;
    color: #000;
  }
  td {
    border-bottom: 1px solid #000;
    background: #fff;
    cursor: pointer;
    word-break: break-all;
    vertical-align: top;
    padding: 10px 0;
  }
`;

export const Th = styled.th`
  background: #ddd;
  font-weight: bold;
  font-size: 12px;
  text-transform: uppercase;
  ${props => props.width && `width: ${props.width}`};
`;
