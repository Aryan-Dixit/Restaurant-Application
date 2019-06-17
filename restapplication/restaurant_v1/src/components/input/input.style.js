import styled from 'styled-components';

export const Input = styled.input`
  width: ${props => (props.block ? '100%' : 'initial')};
  height: 36px;
  padding: 5px 15px;
  font-size: 0.9rem;
  color: #ddd;
  background: #fff;
  border: 1px solid #ddd;
  margin-top: 10px;
`;
